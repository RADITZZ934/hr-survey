package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRateLimit(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(RateLimit())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// The rate limit allows a burst of 30 requests.
	// Let's make 35 rapid requests from the same client IP.
	// The first 30 should succeed (status 200).
	// The next 5 should fail with StatusTooManyRequests (429).
	for i := 0; i < 35; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		// Set remote address to simulate the client IP
		req.RemoteAddr = "192.0.2.1:1234"

		r.ServeHTTP(w, req)

		if i < 30 {
			if w.Code != http.StatusOK {
				t.Fatalf("Expected request %d to succeed, but got status %d", i+1, w.Code)
			}
		} else {
			if w.Code != http.StatusTooManyRequests {
				t.Fatalf("Expected request %d to be rate-limited (429), but got status %d. Response: %s", i+1, w.Code, w.Body.String())
			}
		}
	}
}
