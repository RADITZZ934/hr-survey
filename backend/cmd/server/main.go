package main

import (
	"fmt"
	"net/http"

	"employee-satisfaction-system/backend/config"
	"employee-satisfaction-system/backend/middleware"
	"employee-satisfaction-system/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize environment and configuration
	config.LoadEnv()
	config.LoadConfig()

	// Connect to Database
	config.ConnectDatabase()
	config.AutoMigrate()
	config.SeedDatabase()

	// Initialize Gin router
	r := gin.Default()

	// Use Rate Limiting Middleware
	r.Use(middleware.RateLimit())

	// Setup CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(244)
			return
		}

		c.Next()
	})

	// Register API Routes
	routes.RegisterRoutes(r)

	// Basic health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "Employee Satisfaction System API is running",
		})
	})

	port := config.AppConfig.Port
	fmt.Printf("🚀 Server is running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("❌ Failed to start server: %v\n", err)
	}
}
