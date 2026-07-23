package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

// GetEmployees lists all users
func GetEmployees(c *gin.Context) {
	var users []models.User
	if err := config.DB.Order("id asc").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}

	for i := range users {
		var avgScore float64
		var count int64
		// We match either by user_id or respondent_id (username or email)
		err := config.DB.Table("response_answers").
			Joins("JOIN responses ON response_answers.response_id = responses.id").
			Where("(responses.user_id = ? OR responses.respondent_id = ? OR responses.respondent_id = ?) AND response_answers.score IS NOT NULL", users[i].ID, users[i].Username, users[i].Email).
			Select("COALESCE(AVG(response_answers.score), 0)").
			Row().Scan(&avgScore)

		if err == nil {
			_ = config.DB.Table("response_answers").
				Joins("JOIN responses ON response_answers.response_id = responses.id").
				Where("(responses.user_id = ? OR responses.respondent_id = ? OR responses.respondent_id = ?) AND response_answers.score IS NOT NULL", users[i].ID, users[i].Username, users[i].Email).
				Count(&count)
		}

		if count > 0 && avgScore > 0 {
			formattedAvgScore := float64(int(avgScore*10)) / 10
			percentage := int((avgScore / 5.0) * 100.0)
			users[i].AvgScore = &formattedAvgScore
			users[i].Percentage = &percentage
		}
	}

	c.JSON(http.StatusOK, users)
}

type EmployeeInput struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password"`
	Role       string `json:"role" binding:"required"`
	Department string `json:"department"`
}

// CreateEmployee creates a new employee
func CreateEmployee(c *gin.Context) {
	var input EmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username:   strings.TrimSpace(input.Username),
		Email:      strings.TrimSpace(input.Email),
		Password:   strings.TrimSpace(input.Password),
		Role:       strings.TrimSpace(input.Role),
		Department: strings.TrimSpace(input.Department),
	}

	if user.Password == "" {
		user.Password = "hrd2026" // Default fallback password
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Email already exists"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateEmployee updates employee details
func UpdateEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	var input EmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Username = strings.TrimSpace(input.Username)
	user.Email = strings.TrimSpace(input.Email)
	user.Role = strings.TrimSpace(input.Role)
	user.Department = strings.TrimSpace(input.Department)
	
	if input.Password != "" {
		user.Password = strings.TrimSpace(input.Password)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteEmployee removes employee
func DeleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

// GetCriticalAlerts retrieves unread critical alerts
func GetCriticalAlerts(c *gin.Context) {
	var alerts []models.Alert
	if err := config.DB.Preload("Survey").Order("created_at desc").Limit(20).Find(&alerts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch alerts"})
		return
	}
	c.JSON(http.StatusOK, alerts)
}

// MarkAlertRead marks an alert as read
func MarkAlertRead(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alert ID"})
		return
	}

	if err := config.DB.Model(&models.Alert{}).Where("id = ?", id).Update("is_read", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update alert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alert marked as read"})
}

type DeptSatisfaction struct {
	Department string  `json:"department"`
	AvgScore   float64 `json:"avg_score"`
	Count      int64   `json:"count"`
}

// GetDepartmentSatisfaction returns average satisfaction scores grouped by department
func GetDepartmentSatisfaction(c *gin.Context) {
	var results []DeptSatisfaction
	err := config.DB.Table("responses").
		Select("responses.respondent_dept as department, COALESCE(AVG(response_answers.score), 0) as avg_score, COUNT(DISTINCT responses.id) as count").
		Joins("JOIN response_answers ON response_answers.response_id = responses.id").
		Where("response_answers.score IS NOT NULL AND responses.respondent_dept != ''").
		Group("responses.respondent_dept").
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch department satisfaction"})
		return
	}
	c.JSON(http.StatusOK, results)
}
