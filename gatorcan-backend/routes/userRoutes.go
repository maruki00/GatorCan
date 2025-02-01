package routes

import (
	"gatorcan-backend/controllers"
	"gatorcan-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.POST("", controllers.CreateUser)
	}
}
