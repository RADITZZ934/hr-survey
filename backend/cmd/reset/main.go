package main

import (
	"log"

	"employee-satisfaction-system/backend/config"
)

func main() {
	config.LoadEnv()
	config.LoadConfig()
	config.ConnectDatabase()

	log.Println("⚠️ Wiping database...")
	// Truncate all tables
	err := config.DB.Exec("TRUNCATE users, survey_categories, surveys, responses, response_answers, action_plans, questions, survey_questions, survey_templates, respondents, alerts, activity_logs CASCADE").Error
	if err != nil {
		log.Fatalf("❌ Failed to truncate database: %v", err)
	}
	log.Println("✅ Database tables truncated successfully.")

	log.Println("🌱 Seeding default users and categories...")
	config.SeedDatabase()
	log.Println("🎉 Database reset to clean state successfully!")
}
