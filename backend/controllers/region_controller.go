package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Regency struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// getDataDir returns the absolute path to the backend/data directory
func getDataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "data")
}

func GetProvinces(c *gin.Context) {
	dataPath := filepath.Join(getDataDir(), "provinces.json")
	data, err := os.ReadFile(dataPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load provinces data"})
		return
	}

	var provinces []Province
	if err := json.Unmarshal(data, &provinces); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse provinces data"})
		return
	}

	c.JSON(http.StatusOK, provinces)
}

func GetRegencies(c *gin.Context) {
	provinceID := c.Query("province_id")

	dataPath := filepath.Join(getDataDir(), "regencies.json")
	data, err := os.ReadFile(dataPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load regencies data"})
		return
	}

	var allRegencies []Regency
	if err := json.Unmarshal(data, &allRegencies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse regencies data"})
		return
	}

	// Filter by province_id if provided
	if provinceID != "" {
		var filtered []Regency
		for _, r := range allRegencies {
			if r.ProvinceID == provinceID {
				filtered = append(filtered, r)
			}
		}
		c.JSON(http.StatusOK, filtered)
		return
	}

	c.JSON(http.StatusOK, allRegencies)
}
