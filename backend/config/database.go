package config

import (
	"fmt"
	"log"
	"time"

	"employee-satisfaction-system/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		AppConfig.DBHost,
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBName,
		AppConfig.DBPort,
		AppConfig.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Database ping failed: %v", err)
	}

	DB = db

	log.Println("✅ PostgreSQL Connected Successfully")
}

func AutoMigrate() {
	log.Println("🔄 Running Database Migration...")
	err := DB.AutoMigrate(
		&models.User{},
		&models.SurveyCategory{},
		&models.Question{},
		&models.SurveyTemplate{},
		&models.TemplateQuestion{},
		&models.Survey{},
		&models.SurveyQuestion{},
		&models.Respondent{},
		&models.Response{},
		&models.ResponseAnswer{},
		&models.ActionPlan{},
		&models.ActivityLog{},
		&models.Alert{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Database migrated")
}