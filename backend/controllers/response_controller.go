package controllers

import (
	"net/http"
	"strconv"
	"time"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"
	"employee-satisfaction-system/backend/services"

	"github.com/gin-gonic/gin"
)

type AnswerInput struct {
	QuestionID uint   `json:"question_id" binding:"required"`
	Score      *int   `json:"score"`
	AnswerText string `json:"answer_text"`
}

type CreateResponseInput struct {
	RespondentID       string        `json:"respondent_id"`
	RespondentDept     string        `json:"respondent_dept"`
	RespondentProvince string        `json:"respondent_province"`
	RespondentRegency  string        `json:"respondent_regency"`
	Answers            []AnswerInput `json:"answers" binding:"required"`
}

func CreateResponse(c *gin.Context) {
	surveyIDStr := c.Param("id")
	surveyID, err := strconv.ParseUint(surveyIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid survey ID"})
		return
	}

	var input CreateResponseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Begin Transaction
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var matchedUser models.User
	var userID *uint = nil
	respondentDept := input.RespondentDept

	if input.RespondentID != "" && input.RespondentID != "ANONYMOUS" {
		// Look up user by username or email
		err := config.DB.Where("username = ? OR email = ?", input.RespondentID, input.RespondentID).First(&matchedUser).Error
		if err == nil {
			userID = &matchedUser.ID
			// Set their registered department to ensure consistency
			if matchedUser.Department != "" {
				respondentDept = matchedUser.Department
			}
		}
	}

	// 1. Create Response
	response := models.Response{
		SurveyID:           uint(surveyID),
		UserID:             userID,
		RespondentID:       input.RespondentID,
		RespondentDept:     respondentDept,
		RespondentProvince: input.RespondentProvince,
		RespondentRegency:  input.RespondentRegency,
		SubmittedAt:        time.Now(),
	}

	if err := tx.Create(&response).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record response"})
		return
	}

	// 2. Create Response Answers & Calculate Score
	totalScore := 0
	ratingCount := 0

	for _, ansInput := range input.Answers {
		ans := models.ResponseAnswer{
			ResponseID: response.ID,
			QuestionID: ansInput.QuestionID,
			Score:      ansInput.Score,
			AnswerText: ansInput.AnswerText,
		}

		if ansInput.Score != nil {
			totalScore += *ansInput.Score
			ratingCount++
		}

		if err := tx.Create(&ans).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save response answer"})
			return
		}
	}

	// Commit Transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit response"})
		return
	}

	avgScore := 0.0
	percentage := 0.0
	category := "N/A"

	if ratingCount > 0 {
		avgScore = float64(totalScore) / float64(ratingCount)
		percentage = (avgScore / 5.0) * 100.0

		if avgScore >= 4.5 {
			category = "Sangat Puas"
		} else if avgScore >= 3.5 {
			category = "Puas"
		} else if avgScore >= 2.5 {
			category = "Cukup"
		} else {
			category = "Perlu Perhatian"
		}

		// Log critical score alerts (< 60%) to database & trigger Telegram notifications
		if percentage < 60.0 {
			var survey models.Survey
			config.DB.First(&survey, response.SurveyID)
			surveyTitle := survey.Title
			if surveyTitle == "" {
				surveyTitle = "Survei Kepuasan Kerja"
			}

			respondentName := response.RespondentID
			if respondentName == "" || respondentName == "ANONYMOUS" {
				respondentName = "Anonim"
			}

			alert := models.Alert{
				SurveyID:   response.SurveyID,
				Respondent: respondentName,
				Score:      avgScore,
				Message:    "Skor kepuasan rendah terdeteksi dari pengisian kuesioner.",
			}
			config.DB.Create(&alert)

			// Send external Telegram notification
			services.SendTelegramAlert(avgScore, surveyTitle, respondentName)
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Response submitted successfully",
		"id":      response.ID,
		"data": gin.H{
			"score":        avgScore,
			"max_score":    5.0,
			"percentage":   int(percentage),
			"category":     category,
			"rating_count": ratingCount,
		},
	})
}
