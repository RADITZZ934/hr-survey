package routes

import (
	"employee-satisfaction-system/backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Auth & Admin
		api.POST("/auth/login", controllers.Login)
		api.GET("/admin/backup", controllers.BackupDatabase)

		// Dashboard Stats
		api.GET("/dashboard/stats", controllers.GetDashboardStats)
		api.GET("/dashboard/trends", controllers.GetSurveyTrends)

		// Surveys
		api.GET("/surveys", controllers.GetSurveys)
		api.POST("/surveys", controllers.CreateSurvey)
		api.GET("/surveys/:id/questions", controllers.GetSurveyQuestions)

		// Categories
		api.GET("/categories", controllers.GetCategories)

		// Reports
		api.GET("/surveys/:id/report", controllers.GetSurveyReport)
		api.GET("/surveys/:id/responses", controllers.GetSurveyResponses)
		api.POST("/surveys/:id/responses", controllers.CreateResponse)

		// Action Plans
		api.GET("/action-plans", controllers.GetActionPlans)
		api.POST("/action-plans", controllers.CreateActionPlan)
		api.PUT("/action-plans/:id", controllers.UpdateActionPlan)

		// Employees CRUD
		api.GET("/admin/employees", controllers.GetEmployees)
		api.POST("/admin/employees", controllers.CreateEmployee)
		api.PUT("/admin/employees/:id", controllers.UpdateEmployee)
		api.DELETE("/admin/employees/:id", controllers.DeleteEmployee)
		api.GET("/admin/departments/satisfaction", controllers.GetDepartmentSatisfaction)

		// Critical Alerts
		api.GET("/admin/alerts", controllers.GetCriticalAlerts)
		api.PUT("/admin/alerts/:id/read", controllers.MarkAlertRead)
	}
}
