package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

// GetStores handles GET /api/LIT/store.
// Mengembalikan daftar toko / DC Laskar Buah untuk integrasi survey external.
func GetStores(c *gin.Context) {
	// Auto sync if empty
	var count int64
	config.DB.Model(&models.Store{}).Count(&count)
	if count == 0 {
		syncFromJSON()
	}

	var stores []models.Store
	if err := config.DB.Order("id asc").Find(&stores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stores"})
		return
	}

	c.JSON(http.StatusOK, stores)
}

// SyncStores handles POST /api/LIT/store/sync.
// Menyinkronkan daftar toko dari stores.json dan API Laskar Buah Pusat.
func SyncStores(c *gin.Context) {
	addedFromJSON := syncFromJSON()
	addedFromLive := syncFromLiveAPI()

	var stores []models.Store
	config.DB.Order("id asc").Find(&stores)

	c.JSON(http.StatusOK, gin.H{
		"status":          "success",
		"message":         fmt.Sprintf("Berhasil menyinkronkan data toko (+%d dari JSON, +%d dari API Pusat)", addedFromJSON, addedFromLive),
		"total_stores":    len(stores),
		"added_from_json": addedFromJSON,
		"added_from_live": addedFromLive,
		"data":            stores,
	})
}

func syncFromJSON() int {
	jsonPath := "data/stores.json"
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		jsonPath = filepath.Join("backend", "data", "stores.json")
	}

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return 0
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return 0
	}

	var stores []models.Store
	if err := json.Unmarshal(byteValue, &stores); err != nil {
		return 0
	}

	var added int
	for _, s := range stores {
		var cnt int64
		config.DB.Model(&models.Store{}).Where("kode = ? OR name = ?", s.Kode, s.Name).Count(&cnt)
		if cnt == 0 {
			if err := config.DB.Create(&s).Error; err == nil {
				added++
			}
		}
	}
	return added
}

type LiveAPIResponse struct {
	Status string `json:"status"`
	Data   []struct {
		NamaToko string `json:"nama_toko"`
	} `json:"data"`
}

func syncFromLiveAPI() int {
	fetchURL := "https://api.laskarbuah.com/api/LIT/CetakLaporanKunjunganToko?start_date=2026-01-01&end_date=2026-12-31"
	apiKey := "LB-IT-2026-9tP2mQv7Wz5XkL1nR4jB8sY3fH6gT9dC0xM7vA2wE5nU8pI1oK4rL7qS0zB3mN6"

	req, err := http.NewRequest(http.MethodGet, fetchURL, nil)
	if err != nil {
		log.Printf("⚠️ Failed to create sync request: %v", err)
		return 0
	}

	req.Header.Set("X-Laskar-Career-Key", apiKey)
	req.Header.Set("User-Agent", "Apidog/1.0.0 (https://apidog.com)")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("⚠️ Failed to connect to live API: %v", err)
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("⚠️ Live API returned status %d", resp.StatusCode)
		return 0
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0
	}

	var apiRes LiveAPIResponse
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return 0
	}

	var maxStore models.Store
	config.DB.Order("id desc").First(&maxStore)
	nextID := maxStore.ID

	var added int
	for _, item := range apiRes.Data {
		name := strings.ToUpper(strings.TrimSpace(item.NamaToko))
		if name == "" {
			continue
		}

		var cnt int64
		config.DB.Model(&models.Store{}).Where("UPPER(name) = ?", name).Count(&cnt)
		if cnt == 0 {
			nextID++
			newStore := models.Store{
				ID:         nextID,
				Kode:       fmt.Sprintf("LB%05d", nextID),
				Name:       name,
				AplikasiID: 1,
				Koordinat:  "",
			}
			if err := config.DB.Create(&newStore).Error; err == nil {
				added++
			}
		}
	}

	return added
}
