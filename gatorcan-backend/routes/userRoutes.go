package routes

import (
	"gatorcan-backend/controllers"
	"gatorcan-backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.POST("/add_user", controllers.CreateUser)
	}
}
