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
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthMiddleware()) // Apply JWT authentication middleware
	{
		userGroup.POST("", controllers.CreateUser)
		userGroup.GET("/:username", controllers.GetUserDetails) // For getting user details
		userGroup.DELETE("/:username", controllers.DeleteUser)
		userGroup.PUT("/update", controllers.UpdateUser)
		userGroup.PUT("/update_role", controllers.UpdateRoles)
	}
	router.POST("/login", controllers.Login)
	return router
}

// SetupTestDB initializes an in-memory SQLite database for testing
func SetupTestDB() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{}) // Create schema
	database.DB.Exec("DELETE FROM users")   // Clear users table
}
