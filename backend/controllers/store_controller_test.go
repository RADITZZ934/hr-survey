package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetStores(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup router and route
	r := gin.Default()
	r.GET("/api/LIT/store", GetStores)

	// Create request
	req, _ := http.NewRequest(http.MethodGet, "/api/LIT/store", nil)
	w := httptest.NewRecorder()

	// Perform request
	r.ServeHTTP(w, req)

	// Assertions
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var stores []StoreLocation
	err := json.Unmarshal(w.Body.Bytes(), &stores)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if len(stores) == 0 {
		t.Fatal("Expected non-empty list of stores")
	}

	expectedFirstName := "DC Magetan 1"
	if stores[0].Name != expectedFirstName {
		t.Errorf("Expected first store name to be %s, got %s", expectedFirstName, stores[0].Name)
	}
}
