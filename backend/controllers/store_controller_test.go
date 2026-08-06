package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/models"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// Change working directory to backend root so it can load .env and data/stores.json
	_ = os.Chdir("..")
	config.LoadEnv()
	config.LoadConfig()
	config.ConnectDatabase()
	config.AutoMigrate()
	config.SeedDatabase()

	os.Exit(m.Run())
}

func TestGetStores(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/api/LIT/store", GetStores)

	req, _ := http.NewRequest(http.MethodGet, "/api/LIT/store", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var stores []models.Store
	err := json.Unmarshal(w.Body.Bytes(), &stores)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if len(stores) == 0 {
		t.Fatal("Expected non-empty list of stores")
	}

	expectedFirstName := "PUSAT"
	if stores[0].Name != expectedFirstName {
		t.Errorf("Expected first store name to be %s, got %s", expectedFirstName, stores[0].Name)
	}
}
