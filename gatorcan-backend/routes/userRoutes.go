package routes

import (
	"gatorcan-backend/controllers"
	"gatorcan-backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.POST("/add_user", controllers.CreateUser)
		userGroup.GET("/:username", controllers.GetUserDetails)
		userGroup.DELETE("/:username", controllers.DeleteUser)
		userGroup.PUT("/update", controllers.UpdateUser)
		userGroup.PUT("/update_role", controllers.UpdateRoles)

	}
}
