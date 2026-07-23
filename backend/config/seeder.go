package config

import (
	"log"

	"employee-satisfaction-system/backend/models"
)

func SeedDatabase() {
	log.Println("🔄 Cleaning up old dummy data and seeding essential tables...")

	// WIPE OUT previous dummy data completely from PostgreSQL database (Commented out to prevent data loss on restart)
	// DB.Exec("TRUNCATE surveys, responses, response_answers, action_plans, questions, survey_questions, survey_templates CASCADE")
	// log.Println("⚠️ Cleaned all dummy surveys, questions, responses, and action plans")

	// 1. Seed Users
	var userCount int64
	DB.Model(&models.User{}).Count(&userCount)
	if userCount == 0 {
		users := []models.User{
			{Username: "hradmin", Email: "hradmin@laskarbuah.com", Password: "hrd2026", Role: "admin"},
			{Username: "admin", Email: "admin@company.com", Password: "hrd2026", Role: "admin"},
			{Username: "diana_hr", Email: "diana.r@company.com", Password: "hrd2026", Role: "hr"},
			{Username: "john_manager", Email: "john.d@company.com", Password: "hrd2026", Role: "manager"},
		}
		for _, u := range users {
			DB.Create(&u)
		}
		log.Println("✅ Seeded Users")
	}

	// 2. Seed Survey Categories
	var categoryCount int64
	DB.Model(&models.SurveyCategory{}).Count(&categoryCount)
	if categoryCount == 0 {
		categories := []models.SurveyCategory{
			{Name: "Work-Life Balance", Description: "Work hours, remote flexibility, and workload management."},
			{Name: "Team Collaboration", Description: "Supportive team environment, respect, and communication."},
			{Name: "Manager Support", Description: "Feedback quality, care, and manager communication."},
			{Name: "Compensation & Benefits", Description: "Base salary, health insurance, and bonuses."},
			{Name: "Career Growth", Description: "Training, promotions, and clear growth path."},
		}
		for _, c := range categories {
			DB.Create(&c)
		}
		log.Println("✅ Seeded Survey Categories")
	}
}
