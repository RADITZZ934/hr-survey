package controllers

import (
	"net/http"
	"time"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

func GetDashboardStats(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	var totalResponses int64
	var actionPlansCount int64
	var avgScore float64

	dbResponse := config.DB.Model(&models.Response{})
	dbActionPlan := config.DB.Model(&models.ActionPlan{})
	dbAnswer := config.DB.Table("response_answers")

	if startDate != "" {
		dbResponse = dbResponse.Where("submitted_at >= ?", startDate+" 00:00:00")
		dbActionPlan = dbActionPlan.Where("created_at >= ?", startDate+" 00:00:00")
		dbAnswer = dbAnswer.Joins("JOIN responses ON response_answers.response_id = responses.id").
			Where("responses.submitted_at >= ?", startDate+" 00:00:00")
	}
	if endDate != "" {
		dbResponse = dbResponse.Where("submitted_at <= ?", endDate+" 23:59:59")
		dbActionPlan = dbActionPlan.Where("created_at <= ?", endDate+" 23:59:59")
		if startDate == "" {
			dbAnswer = dbAnswer.Joins("JOIN responses ON response_answers.response_id = responses.id")
		}
		dbAnswer = dbAnswer.Where("responses.submitted_at <= ?", endDate+" 23:59:59")
	}

	// Count responses
	if err := dbResponse.Count(&totalResponses).Error; err != nil {
		totalResponses = 0
	}

	// Count action plans
	if err := dbActionPlan.Count(&actionPlansCount).Error; err != nil {
		actionPlansCount = 0
	}

	// Compute average score
	row := dbAnswer.Where("response_answers.score IS NOT NULL").Select("COALESCE(AVG(response_answers.score), 0)").Row()
	_ = row.Scan(&avgScore)

	// Format avgScore to 1 decimal place
	formattedAvgScore := float64(int(avgScore*10)) / 10

	// Calculate completion rate based on a static target of 100 respondents
	targetRespondents := int64(100)

	completionRate := float64(totalResponses) / float64(targetRespondents) * 100
	if completionRate > 100.0 {
		completionRate = 100.0
	}
	completionRate = float64(int(completionRate*10)) / 10

	c.JSON(http.StatusOK, gin.H{
		"avgScore":         formattedAvgScore,
		"totalResponses":   totalResponses,
		"completionRate":   completionRate,
		"actionPlansCount": actionPlansCount,
	})
}

func GetSurveyTrends(c *gin.Context) {
	startDateStr := c.Query("startDate")
	endDateStr := c.Query("endDate")

	// 1. Get 3 newest surveys
	var surveys []models.Survey
	if err := config.DB.Order("created_at desc").Limit(3).Find(&surveys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch surveys"})
		return
	}

	// 2. Generate list of dates based on start/end dates
	var dateStrings []string
	var labels []string

	if startDateStr != "" && endDateStr != "" {
		start, err1 := time.Parse("2006-01-02", startDateStr)
		end, err2 := time.Parse("2006-01-02", endDateStr)
		if err1 == nil && err2 == nil {
			days := int(end.Sub(start).Hours()/24) + 1
			if days > 31 {
				days = 31
				start = end.AddDate(0, 0, -30)
			}
			for i := 0; i < days; i++ {
				t := start.AddDate(0, 0, i)
				dateStrings = append(dateStrings, t.Format("2006-01-02"))
				labels = append(labels, t.Format("Jan 02"))
			}
		}
	}

	if len(dateStrings) == 0 {
		for i := 6; i >= 0; i-- {
			t := time.Now().AddDate(0, 0, -i)
			dateStrings = append(dateStrings, t.Format("2006-01-02"))
			labels = append(labels, t.Format("Jan 02"))
		}
	}

	// 3. For each survey, get response count on each date
	type TrendSeries struct {
		ID    uint    `json:"id"`
		Title string  `json:"title"`
		Data  []int64 `json:"data"`
	}

	datasets := []TrendSeries{}

	for _, survey := range surveys {
		counts := make([]int64, len(dateStrings))
		for idx, dateStr := range dateStrings {
			var count int64
			startOfDay := dateStr + " 00:00:00"
			endOfDay := dateStr + " 23:59:59"
			_ = config.DB.Model(&models.Response{}).
				Where("survey_id = ? AND submitted_at >= ? AND submitted_at <= ?", survey.ID, startOfDay, endOfDay).
				Count(&count)
			counts[idx] = count
		}
		datasets = append(datasets, TrendSeries{
			ID:    survey.ID,
			Title: survey.Title,
			Data:  counts,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"labels":   labels,
		"datasets": datasets,
	})
}


