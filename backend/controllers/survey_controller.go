package controllers

import (
	"net/http"
	"time"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

func GetSurveys(c *gin.Context) {
	var surveys []models.Survey
	if err := config.DB.Order("created_at desc").Find(&surveys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch surveys"})
		return
	}

	for i := range surveys {
		var count int64
		config.DB.Model(&models.Response{}).Where("survey_id = ?", surveys[i].ID).Count(&count)
		surveys[i].ResponsesCount = count
	}

	c.JSON(http.StatusOK, surveys)
}

type QuestionInput struct {
	Text       string `json:"text" binding:"required"`
	Type       string `json:"type" binding:"required"`
	IsRequired bool   `json:"is_required"`
}

type CreateSurveyInput struct {
	Title       string          `json:"title" binding:"required"`
	Description string          `json:"description"`
	StartDate   string          `json:"start_date" binding:"required"`
	EndDate     string          `json:"end_date" binding:"required"`
	TemplateID  *uint           `json:"template_id,omitempty"`
	Questions   []QuestionInput `json:"questions"`
}

func CreateSurvey(c *gin.Context) {
	var input CreateSurveyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	end, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	// Begin GORM Transaction
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	survey := models.Survey{
		Title:       input.Title,
		Description: input.Description,
		StartDate:   start,
		EndDate:     end,
		TemplateID:  input.TemplateID,
		Status:      "active",
		CreatedBy:   1, // Admin default creator
	}

	if err := tx.Create(&survey).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create survey"})
		return
	}

	// Create questions and associations
	for _, qInput := range input.Questions {
		questionType := "scale"
		if qInput.Type == "essay" || qInput.Type == "text" {
			questionType = "text"
		}

		categoryID := uint(1) // Category 1 for Rating Bintang
		if questionType == "text" {
			categoryID = 2 // Category 2 for Essay
		}

		question := models.Question{
			Text:       qInput.Text,
			Type:       questionType,
			CategoryID: categoryID,
			IsRequired: qInput.IsRequired,
		}

		if err := tx.Create(&question).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create questions"})
			return
		}

		association := models.SurveyQuestion{
			SurveyID:   survey.ID,
			QuestionID: question.ID,
		}

		if err := tx.Create(&association).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to link questions"})
			return
		}
	}

	// Commit Transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction commit failed"})
		return
	}

	c.JSON(http.StatusCreated, survey)
}

func GetSurveyQuestions(c *gin.Context) {
	surveyID := c.Param("id")
	
	var questions []models.Question
	err := config.DB.
		Joins("JOIN survey_questions ON survey_questions.question_id = questions.id").
		Where("survey_questions.survey_id = ?", surveyID).
		Preload("Category").
		Find(&questions).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch survey questions"})
		return
	}

	c.JSON(http.StatusOK, questions)
}
