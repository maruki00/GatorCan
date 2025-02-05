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
	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware()) // Apply JWT authentication middleware
	{
		userGroup.POST("", controllers.CreateUser)
		userGroup.GET("/:username", controllers.GetUserDetails) // For getting user details
	}
	return router
}

// SetupTestDB initializes an in-memory SQLite database for testing
func SetupTestDB() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{}) // Create schema
}
