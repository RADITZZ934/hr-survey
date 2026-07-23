package controllers

import (
	"net/http"
	"strconv"
	"time"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

type ActionPlanResponse struct {
	ID               uint   `json:"id"`
	SurveyID         uint   `json:"surveyId"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Status           string `json:"status"`
	TargetDate       string `json:"targetDate"`
	AssigneeName     string `json:"assigneeName"`
	AssigneeInitials string `json:"assigneeInitials"`
}

func GetActionPlans(c *gin.Context) {
	var plans []models.ActionPlan
	query := config.DB.Preload("Assignee")

	surveyIDStr := c.Query("survey_id")
	if surveyIDStr != "" {
		if surveyID, err := strconv.ParseUint(surveyIDStr, 10, 32); err == nil {
			query = query.Where("survey_id = ?", uint(surveyID))
		}
	}

	if err := query.Order("created_at desc").Find(&plans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch action plans"})
		return
	}

	var responseList []ActionPlanResponse
	for _, p := range plans {
		assigneeName := "Diana R."
		assigneeInitials := "DR"
		if p.Assignee != nil {
			assigneeName = p.Assignee.Username
			assigneeInitials = string(assigneeName[0]) + string(assigneeName[len(assigneeName)-1])
		}

		targetDateStr := ""
		if p.TargetDate != nil {
			targetDateStr = p.TargetDate.Format("2006-01-02")
		}

		responseList = append(responseList, ActionPlanResponse{
			ID:               p.ID,
			SurveyID:         p.SurveyID,
			Title:            p.Title,
			Description:      p.Description,
			Status:           p.Status,
			TargetDate:       targetDateStr,
			AssigneeName:     assigneeName,
			AssigneeInitials: assigneeInitials,
		})
	}

	c.JSON(http.StatusOK, responseList)
}

type CreatePlanInput struct {
	SurveyID    uint   `json:"surveyId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TargetDate  string `json:"targetDate"`
	AssigneeID  *uint  `json:"assigneeId,omitempty"`
}

func CreateActionPlan(c *gin.Context) {
	var input CreatePlanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var targetTime *time.Time
	if input.TargetDate != "" {
		if t, err := time.Parse("2006-01-02", input.TargetDate); err == nil {
			targetTime = &t
		}
	}

	// Default assignee to Diana HR if not provided
	var assigneeID uint = 2
	if input.AssigneeID != nil {
		assigneeID = *input.AssigneeID
	}

	plan := models.ActionPlan{
		SurveyID:    input.SurveyID,
		Title:       input.Title,
		Description: input.Description,
		Status:      "pending",
		TargetDate:  targetTime,
		AssigneeID:  &assigneeID,
		CreatedBy:   1, // Admin default creator
	}

	if err := config.DB.Create(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create action plan"})
		return
	}

	// Preload assignee to return formatted response
	config.DB.Preload("Assignee").First(&plan, plan.ID)

	assigneeName := "Diana R."
	assigneeInitials := "DR"
	if plan.Assignee != nil {
		assigneeName = plan.Assignee.Username
		assigneeInitials = string(assigneeName[0]) + string(assigneeName[len(assigneeName)-1])
	}

	targetDateStr := ""
	if plan.TargetDate != nil {
		targetDateStr = plan.TargetDate.Format("2006-01-02")
	}

	c.JSON(http.StatusCreated, ActionPlanResponse{
		ID:               plan.ID,
		SurveyID:         plan.SurveyID,
		Title:            plan.Title,
		Description:      plan.Description,
		Status:           plan.Status,
		TargetDate:       targetDateStr,
		AssigneeName:     assigneeName,
		AssigneeInitials: assigneeInitials,
	})
}

type UpdatePlanInput struct {
	Status string `json:"status"`
}

func UpdateActionPlan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action plan ID"})
		return
	}

	var input UpdatePlanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var plan models.ActionPlan
	if err := config.DB.First(&plan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Action plan not found"})
		return
	}

	plan.Status = input.Status
	if err := config.DB.Save(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update action plan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Action plan updated successfully"})
}
