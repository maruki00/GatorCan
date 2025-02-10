package tests

import (
	"gatorcan-backend/controllers"
	"gatorcan-backend/database"
	"gatorcan-backend/middleware"
	"gatorcan-backend/models"

	"github.com/gin-gonic/gin"
)

// SetupTestRouter initializes a test Gin router
func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", controllers.Login)

	// Admin-only Routes
	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.AuthMiddleware(string(models.Admin)))
	{
		adminGroup.POST("/add_user", controllers.CreateUser)
		adminGroup.DELETE("/:username", controllers.DeleteUser)
		adminGroup.PUT("/update_role", controllers.UpdateRoles)

	}
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthMiddleware(string(models.Student)))
	{
		userGroup.GET("/:username", controllers.GetUserDetails)
		userGroup.PUT("/update", controllers.UpdateUser)

	}

	// Instructor-only Routes
	instructorRoutes := router.Group("/instructor")
	instructorRoutes.Use(middleware.AuthMiddleware(string(models.Instructor)))
	{
		//instructorRoutes.POST("/upload-assignment", UploadAssignmentHandler)
	}
	return router
}

// SetupTestDB initializes an in-memory SQLite database for testing
func SetupTestDB() {

	database.Connect()
	database.DB.AutoMigrate(&models.User{}) // Create schema
	database.DB.Exec("insert into roles (created_at, updated_At, name) values(datetime('now'),datetime('now'),'student');")
	database.DB.Exec("insert into roles (created_at, updated_At, name) values(datetime('now'),datetime('now'),'admin');")
	database.DB.Exec("insert into roles (created_at, updated_At, name) values(datetime('now'),datetime('now'),'instructor');")
	database.DB.Exec("insert into roles (created_at, updated_At, name) values(datetime('now'),datetime('now'),'teaching_assistant');")
	database.DB.Exec("DELETE FROM users") // Clear users table
	database.DB.Exec("DELETE FROM roles") // Clear roles table
}
