package tests

import (
	"fmt"
	"gatorcan-backend/controllers"
	"gatorcan-backend/database"
	"gatorcan-backend/middleware"
	"gatorcan-backend/models"
	"time"

	"gatorcan-backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SetupTestRouter initializes a test Gin router
func SetupTestRouter() *gin.Engine {
	logger := utils.Log()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
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
		adminGroup.DELETE("/update_role", func(c *gin.Context) {
			controllers.UpdateRoles(c, logger)
		})

	}
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthMiddleware(logger, string(models.Student)))
	{
		userGroup.GET("/:username", func(c *gin.Context) {
			controllers.GetUserDetails(c, logger)
		})
		userGroup.GET("/update", func(c *gin.Context) {
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
	for i := 1; i <= 30; i++ {
		capacity := 50 + i // Example: Capacity varies between 51-80
		enrollmentCount := 0

		course := models.Course{
			Name:        fmt.Sprintf("Course %d", i),
			Description: fmt.Sprintf("Description for course %d", i),
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(0, 1, 0), // Ends in 1 month
			Capacity:    capacity,
			Enrolled:    enrollmentCount,
		}
		database.DB.Create(&course)
	}

	studentRole := models.Role{Name: "student"}
	adminRole := models.Role{Name: "admin"}
	if err := database.DB.Create(&studentRole).Error; err != nil {
		fmt.Printf("Failed to create 'student' role: %v", err)
	}
	if err := database.DB.Create(&adminRole).Error; err != nil {
		fmt.Printf("Failed to create 'admin' role: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Failed to hash password: %v", err)
	}

	testUser := models.User{
		Username: "teststudent",
		Email:    "teststudent@example.com",
		Password: string(hashedPassword),
		Roles:    []*models.Role{&studentRole}, // Assign "student" role
	}
	if err := database.DB.Create(&testUser).Error; err != nil {
		fmt.Printf("Failed to create test user: %v\n", err)

	}
}
