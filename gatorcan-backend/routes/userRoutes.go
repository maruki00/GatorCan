package routes

import (
	"gatorcan-backend/controllers"
	"gatorcan-backend/middleware"
	"gatorcan-backend/models"
	"log"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, logger *log.Logger) {

	//  Public Routes
	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, logger)
	})

	// Admin-only Routes
	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.AuthMiddleware(logger, string(models.Admin)))
	{
		adminGroup.POST("/add_user", func(c *gin.Context) {
			controllers.CreateUser(c, logger)
		})
		adminGroup.DELETE("/:username", func(c *gin.Context) {
			controllers.DeleteUser(c, logger)
		})
		adminGroup.POST("/update_role", func(c *gin.Context) {
			controllers.UpdateRoles(c, logger)
		})

	}
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthMiddleware(logger, string(models.Student), string(models.Admin)))
	{
		userGroup.GET("/:username", func(c *gin.Context) {
			controllers.GetUserDetails(c, logger)
		})
		userGroup.PUT("/update", func(c *gin.Context) {
			controllers.UpdateUser(c, logger)
		})

	}

	// Instructor-only Routes
	instructorRoutes := router.Group("/instructor")
	instructorRoutes.Use(middleware.AuthMiddleware(logger, string(models.Instructor)))
	{
		//instructorRoutes.POST("/upload-assignment", UploadAssignmentHandler)
	}

	courseGroup := router.Group("/courses")
	courseGroup.Use(middleware.AuthMiddleware(logger, string(models.Student)))
	{
		courseGroup.GET("/enrolled", func(c *gin.Context) {
			controllers.GetEnrolledCourses(c, logger)
		})
		//courseGroup.POST("/enroll", controllers.EnrollCourse)
		courseGroup.GET("/", func(c *gin.Context) {
			controllers.GetCourses(c, logger)
		})

		courseGroup.POST("/enroll", func(c *gin.Context) {
			controllers.EnrollInCourse(c, logger)
		})
	}
}
