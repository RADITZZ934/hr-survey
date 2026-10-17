package controllers

import (
	"net/http"
	"strings"
	"time"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

// SubmitSurveyResponseInput mendefinisikan payload JSON yang diterima dari frontend.
type SubmitSurveyResponseInput struct {
	SurveyID       uint   `json:"survey_id" binding:"required"`
	IDStore        string `json:"id_store" binding:"required"`
	NamaResponden  string `json:"nama_responden" binding:"required"`
	EmailResponden string `json:"email_responden"`
	Penilaian      string `json:"penilaian" binding:"required"`
}

// SubmitSurveyResponse handles POST /api/surveys/submit.
// Menerima data hasil survey dari customer via QR code toko dan menyimpannya ke tabel survey_responses.
func SubmitSurveyResponse(c *gin.Context) {
	var input SubmitSurveyResponseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Data tidak valid: " + err.Error(),
		})
		return
	}

	clientIP := c.ClientIP()
	isFlagged := false
	flagReason := ""
	cleanEmail := strings.ToLower(strings.TrimSpace(input.EmailResponden))

	if cleanEmail != "" {
		if strings.HasSuffix(cleanEmail, "@laskarbuah.com") ||
			strings.HasSuffix(cleanEmail, "@bazzar.com") ||
			strings.HasSuffix(cleanEmail, "@holding.com") {
			isFlagged = true
			flagReason = "Internal Staff Email"
		} else {
			var existingCount int64
			sevenDaysAgo := time.Now().AddDate(0, 0, -7)
			config.DB.Model(&models.SurveyResponse{}).
				Where("email_responden = ? AND id_store = ? AND created_at >= ?", cleanEmail, input.IDStore, sevenDaysAgo).
				Count(&existingCount)

			if existingCount > 0 {
				isFlagged = true
				flagReason = "Duplicate Submission within 7 days"
			}
		}
	}

	surveyResponse := models.SurveyResponse{
		SurveyID:       input.SurveyID,
		IDStore:        input.IDStore,
		NamaResponden:  input.NamaResponden,
		EmailResponden: cleanEmail,
		Penilaian:      input.Penilaian,
		IPAddress:      clientIP,
		IsFlagged:      isFlagged,
		FlagReason:     flagReason,
	}

	if err := config.DB.Create(&surveyResponse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal menyimpan data: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Data berhasil disimpan",
	})
}
