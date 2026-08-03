package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

type CategoryScore struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

func GetSurveyReport(c *gin.Context) {
	surveyIDStr := c.Param("id")
	surveyID, err := strconv.ParseUint(surveyIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid survey ID"})
		return
	}

	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	// 1. Fetch Survey metadata
	var survey models.Survey
	if err := config.DB.First(&survey, surveyID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Survey not found"})
		return
	}

	// 2. Count responses with date filter
	var totalResponses int64
	dbResponse := config.DB.Model(&models.Response{}).Where("survey_id = ?", surveyID)
	if startDate != "" {
		dbResponse = dbResponse.Where("submitted_at >= ?", startDate+" 00:00:00")
	}
	if endDate != "" {
		dbResponse = dbResponse.Where("submitted_at <= ?", endDate+" 23:59:59")
	}
	_ = dbResponse.Count(&totalResponses)

	// 3. Compute overall average score for this survey with date filter
	var avgScore float64
	dbAnswer := config.DB.Table("response_answers").
		Joins("JOIN responses ON response_answers.response_id = responses.id").
		Where("responses.survey_id = ? AND response_answers.score IS NOT NULL", surveyID)
	if startDate != "" {
		dbAnswer = dbAnswer.Where("responses.submitted_at >= ?", startDate+" 00:00:00")
	}
	if endDate != "" {
		dbAnswer = dbAnswer.Where("responses.submitted_at <= ?", endDate+" 23:59:59")
	}
	row := dbAnswer.Select("COALESCE(AVG(response_answers.score), 0)").Row()
	_ = row.Scan(&avgScore)
	formattedAvgScore := float64(int(avgScore*10)) / 10

	// 4. Compute category averages with date filter
	var categories []CategoryScore
	dbCat := config.DB.Table("response_answers").
		Joins("JOIN responses ON response_answers.response_id = responses.id").
		Joins("JOIN questions ON response_answers.question_id = questions.id").
		Joins("JOIN survey_categories ON questions.category_id = survey_categories.id").
		Where("responses.survey_id = ? AND response_answers.score IS NOT NULL", surveyID)
	if startDate != "" {
		dbCat = dbCat.Where("responses.submitted_at >= ?", startDate+" 00:00:00")
	}
	if endDate != "" {
		dbCat = dbCat.Where("responses.submitted_at <= ?", endDate+" 23:59:59")
	}

	rows, err := dbCat.Select("survey_categories.name, AVG(response_answers.score)").
		Group("survey_categories.name").
		Rows()

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var cat CategoryScore
			if err := rows.Scan(&cat.Name, &cat.Score); err == nil {
				// Round score to 1 decimal place
				cat.Score = float64(int(cat.Score*10)) / 10
				categories = append(categories, cat)
			}
		}
	}

	// Fallback to seeded categories if no answers yet
	if len(categories) == 0 {
		categories = []CategoryScore{
			{Name: "Work-Life Balance", Score: 0},
			{Name: "Team Collaboration", Score: 0},
			{Name: "Manager Support", Score: 0},
			{Name: "Compensation & Benefits", Score: 0},
			{Name: "Career Growth", Score: 0},
		}
	}

	// 5. Predefined strengths and improvements for demonstration
	strengths := "Work-Life Balance and Team Collaboration scored exceptionally high. Team members appreciate remote work flexibilities and supportive management styles."
	improvements := "Compensation transparency and career path clarity remain top concerns. Focus on introducing development frameworks and reviewing bonus schemas."

	if surveyID == 2 {
		strengths = "Workplace safety and physical environments are highly rated. Employees feel the office is comfortable and supportive."
		improvements = "Mental wellness resources are lacking, and team workloads lead to high stress levels. Need to introduce wellness sessions."
	} else if surveyID == 3 {
		strengths = "Basic health insurance coverage is highly appreciated. Bonuses are perceived as fair when targets are met."
		improvements = "Base salary competitiveness is rated low compared to market standards. Wellness allowances are requested."
	}

	// Count action plans generated for this survey
	var actionPlansCount int64
	_ = config.DB.Model(&models.ActionPlan{}).Where("survey_id = ?", surveyID).Count(&actionPlansCount)

	// Calculate completion rate based on a static target of 100 respondents
	targetRespondents := int64(100)

	completionRate := float64(totalResponses) / float64(targetRespondents) * 100
	if completionRate > 100.0 {
		completionRate = 100.0
	}
	completionRate = float64(int(completionRate*10)) / 10

	// 6. Compute rating distribution (count per star 1-5) for accumulation chart
	ratingDistribution := map[string]int64{
		"1": 0, "2": 0, "3": 0, "4": 0, "5": 0,
	}

	dbDist := config.DB.Table("response_answers").
		Joins("JOIN responses ON response_answers.response_id = responses.id").
		Where("responses.survey_id = ? AND response_answers.score IS NOT NULL", surveyID)
	if startDate != "" {
		dbDist = dbDist.Where("responses.submitted_at >= ?", startDate+" 00:00:00")
	}
	if endDate != "" {
		dbDist = dbDist.Where("responses.submitted_at <= ?", endDate+" 23:59:59")
	}

	distRows, err := dbDist.Select("response_answers.score, COUNT(*) as cnt").
		Group("response_answers.score").
		Rows()

	if err == nil {
		defer distRows.Close()
		for distRows.Next() {
			var score int
			var cnt int64
			if err := distRows.Scan(&score, &cnt); err == nil {
				key := strconv.Itoa(score)
				if _, ok := ratingDistribution[key]; ok {
					ratingDistribution[key] = cnt
				}
			}
		}
	}

	// 7. Compute satisfaction categories count
	satisfactionCategories := map[string]int64{
		"Sangat Puas":     0,
		"Puas":            0,
		"Cukup":           0,
		"Perlu Perhatian": 0,
	}

	// Query per-response average scores to categorize each respondent
	catRows, catErr := config.DB.Table("response_answers").
		Select("response_answers.response_id, AVG(response_answers.score) as avg_score").
		Joins("JOIN responses ON response_answers.response_id = responses.id").
		Where("responses.survey_id = ? AND response_answers.score IS NOT NULL", surveyID).
		Group("response_answers.response_id").
		Rows()

	if catErr == nil {
		defer catRows.Close()
		for catRows.Next() {
			var responseID uint
			var respAvg float64
			if err := catRows.Scan(&responseID, &respAvg); err == nil {
				if respAvg >= 4.5 {
					satisfactionCategories["Sangat Puas"]++
				} else if respAvg >= 3.5 {
					satisfactionCategories["Puas"]++
				} else if respAvg >= 2.5 {
					satisfactionCategories["Cukup"]++
				} else {
					satisfactionCategories["Perlu Perhatian"]++
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"avgScore":                formattedAvgScore,
		"totalResponses":          totalResponses,
		"completionRate":          completionRate,
		"actionPlansCount":        actionPlansCount,
		"strengths":               strengths,
		"improvements":            improvements,
		"categories":              categories,
		"ratingDistribution":      ratingDistribution,
		"satisfactionCategories":  satisfactionCategories,
	})
}

type AnswerDetail struct {
	Question string `json:"question"`
	Score    *int   `json:"score,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

type RespondentLog struct {
	ID                 uint           `json:"id"`
	Initials           string         `json:"initials"`
	Name               string         `json:"name"`
	Email              string         `json:"email"`
	Department         string         `json:"department"`
	RespondentProvince string         `json:"respondent_province"`
	RespondentRegency  string         `json:"respondent_regency"`
	AvgRating          float64        `json:"avgRating"`
	SubmittedAt        string         `json:"submittedAt"`
	Answers            []AnswerDetail `json:"answers"`
}

func GetSurveyResponses(c *gin.Context) {
	surveyIDStr := c.Param("id")
	surveyID, err := strconv.ParseUint(surveyIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid survey ID"})
		return
	}

	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	// Get page and limit from query params, setting defaults (limit = 30)
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "30")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 30
	}

	offset := (page - 1) * limit

	// Create count query to get total matching responses
	var totalCount int64
	countQuery := config.DB.Where("survey_id = ?", surveyID)
	if startDate != "" {
		countQuery = countQuery.Where("submitted_at >= ?", startDate+" 00:00:00")
	}
	if endDate != "" {
		countQuery = countQuery.Where("submitted_at <= ?", endDate+" 23:59:59")
	}
	if err := countQuery.Model(&models.Response{}).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count responses"})
		return
	}

	// Create select query to get responses with pagination
	var responses []models.Response
	findQuery := config.DB.Where("survey_id = ?", surveyID)
	if startDate != "" {
		findQuery = findQuery.Where("submitted_at >= ?", startDate+" 00:00:00")
	}
	if endDate != "" {
		findQuery = findQuery.Where("submitted_at <= ?", endDate+" 23:59:59")
	}
	if err := findQuery.Order("submitted_at desc").Limit(limit).Offset(offset).Find(&responses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch responses"})
		return
	}

	var logList []RespondentLog
	for _, r := range responses {
		// Fetch answers
		var answers []models.ResponseAnswer
		if err := config.DB.Where("response_id = ?", r.ID).Preload("Question").Find(&answers).Error; err != nil {
			continue
		}

		var answerDetails []AnswerDetail
		var totalScore int
		var scoreCount int

		for _, ans := range answers {
			detail := AnswerDetail{
				Question: ans.Question.Text,
			}
			if ans.Score != nil {
				detail.Score = ans.Score
				totalScore += *ans.Score
				scoreCount++
			}
			if ans.AnswerText != "" {
				detail.Answer = ans.AnswerText
			}
			answerDetails = append(answerDetails, detail)
		}

		var avgRating float64
		if scoreCount > 0 {
			avgRating = float64(totalScore) / float64(scoreCount)
		}

		// Determine respondent name
		name := "Anonim"
		email := ""
		initials := "AN"

		if r.RespondentID != "" && r.RespondentID != "ANONYMOUS" {
			name = r.RespondentID
			parts := strings.Fields(name)
			if len(parts) >= 2 {
				initials = strings.ToUpper(string(parts[0][0]) + string(parts[1][0]))
			} else if len(parts) == 1 && len(parts[0]) > 0 {
				initials = strings.ToUpper(string(parts[0][0]))
			}
			email = strings.ToLower(strings.ReplaceAll(name, " ", ".")) + "@company.com"
		}

		dept := r.RespondentDept
		if dept == "ANONYMOUS" || dept == "" {
			dept = "-"
		}

		logList = append(logList, RespondentLog{
			ID:                 r.ID,
			Initials:           initials,
			Name:               name,
			Email:              email,
			Department:         dept,
			RespondentProvince: r.RespondentProvince,
			RespondentRegency:  r.RespondentRegency,
			AvgRating:          avgRating,
			SubmittedAt:        r.SubmittedAt.Format(time.RFC1123),
			Answers:            answerDetails,
		})
	}

	// Calculate total pages
	totalPages := int(totalCount) / limit
	if int(totalCount)%limit != 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       logList,
		"total":      totalCount,
		"page":       page,
		"limit":      limit,
		"totalPages": totalPages,
	})
}
