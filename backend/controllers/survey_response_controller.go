package controllers

import (
	"net/http"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

// SubmitSurveyResponseInput mendefinisikan payload JSON yang diterima dari frontend.
type SubmitSurveyResponseInput struct {
	SurveyID      uint   `json:"survey_id" binding:"required"`
	IDStore       string `json:"id_store" binding:"required"`
	NamaResponden string `json:"nama_responden" binding:"required"`
	Penilaian     string `json:"penilaian" binding:"required"`
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

	surveyResponse := models.SurveyResponse{
		SurveyID:      input.SurveyID,
		IDStore:       input.IDStore,
		NamaResponden: input.NamaResponden,
		Penilaian:     input.Penilaian,
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
