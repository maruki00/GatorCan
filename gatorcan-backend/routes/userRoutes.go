package routes

import (
	"gatorcan-backend/controllers"
	"gatorcan-backend/middleware"
	"gatorcan-backend/models"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	//  Public Routes
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
}
