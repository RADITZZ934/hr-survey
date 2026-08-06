package controllers

import (
	"net/http"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

// GetStores handles GET /api/LIT/store.
// Mengembalikan daftar toko / DC Laskar Buah untuk integrasi survey external.
func GetStores(c *gin.Context) {
	var stores []models.Store
	if err := config.DB.Order("id asc").Find(&stores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stores"})
		return
	}

	c.JSON(http.StatusOK, stores)
}
