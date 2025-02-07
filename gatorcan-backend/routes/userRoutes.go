package routes

import (
	"gatorcan-backend/controllers"
	"gatorcan-backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	//  Public Routes
	router.POST("/login", controllers.Login)

	// Admin-only Routes
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthMiddleware("admin"))
	{
		userGroup.POST("/add_user", controllers.CreateUser)
		userGroup.GET("/:username", controllers.GetUserDetails)
		userGroup.DELETE("/:username", controllers.DeleteUser)
		userGroup.PUT("/update", controllers.UpdateUser)
		userGroup.PUT("/update_role", controllers.UpdateRoles)

	}

	// Instructor-only Routes
	instructorRoutes := router.Group("/instructor")
	instructorRoutes.Use(middleware.AuthMiddleware("instructor"))
	{
		//instructorRoutes.POST("/upload-assignment", UploadAssignmentHandler)
	}
}
