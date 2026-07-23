package controllers

import (
	"fmt"
	"net/http"
	"time"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

type DatabaseBackupDump struct {
	App         string                  `json:"app"`
	ExportedAt  time.Time               `json:"exported_at"`
	Users       []models.User           `json:"users"`
	Categories  []models.SurveyCategory `json:"categories"`
	Questions   []models.Question       `json:"questions"`
	Surveys     []models.Survey         `json:"surveys"`
	Respondents []models.Respondent     `json:"respondents"`
	Responses   []models.Response       `json:"responses"`
	Answers     []models.ResponseAnswer `json:"response_answers"`
	ActionPlans []models.ActionPlan     `json:"action_plans"`
}

func BackupDatabase(c *gin.Context) {
	var backup DatabaseBackupDump
	backup.App = "HR SURVEY TOOLS | LASKAR BUAH"
	backup.ExportedAt = time.Now()

	config.DB.Find(&backup.Users)
	config.DB.Find(&backup.Categories)
	config.DB.Find(&backup.Questions)
	config.DB.Find(&backup.Surveys)
	config.DB.Find(&backup.Respondents)
	config.DB.Find(&backup.Responses)
	config.DB.Find(&backup.Answers)
	config.DB.Find(&backup.ActionPlans)

	fileName := fmt.Sprintf("hr_survey_backup_%s.json", time.Now().Format("2006-01-02_150405"))

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, backup)
}
