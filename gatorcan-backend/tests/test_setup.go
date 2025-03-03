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
	userGroup.Use(middleware.AuthMiddleware(logger, string(models.Student), string(models.Admin)))
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
	database.DB.AutoMigrate(&models.User{}, &models.Course{}, &models.Enrollment{}, &models.ActiveCourse{}, &models.Role{}) // Create schema
	database.DB.Exec("insert into roles (created_at, updated_at, name)values(datetime(),datetime(),'student');")
	database.DB.Exec("insert into roles (created_at, updated_at, name)values(datetime(),datetime(),'admin');")
	database.DB.Exec("insert into roles (created_at, updated_at, name)values(datetime(),datetime(),'instructor');")
	database.DB.Exec("insert into users (created_at, updated_at, username, email, password) values(datetime(),datetime(),'instructor', 'instructor@admin.com', '$2y$10$StNLKLEww2O7qiArA/BCmu4gf4RKht6rq19y91YLHcMSlSCv7uGbm');")
	database.DB.Exec("insert into users (created_at, updated_at, username, email, password) values(datetime(),datetime(),'student', 'student@admin.com', '$2y$10$StNLKLEww2O7qiArA/BCmu4gf4RKht6rq19y91YLHcMSlSCv7uGbm');")
	database.DB.Exec("insert into users (created_at, updated_at, username, email, password) values(datetime(),datetime(),'admin', 'admin@admin.com', '$2y$10$StNLKLEww2O7qiArA/BCmu4gf4RKht6rq19y91YLHcMSlSCv7uGbm');")
	database.DB.Exec("insert into user_roles (role_id, user_id)values(1,2);")
	database.DB.Exec("insert into user_roles (role_id, user_id)values(2,3);")
	database.DB.Exec("insert into user_roles (role_id, user_id)values(3,1);")
	database.DB.Exec("insert into courses (created_at, updated_at, name, description)values(datetime(),datetime(),'ADS', 'a course');")
	database.DB.Exec("insert into courses (created_at, updated_at, name, description)values(datetime(),datetime(),'Data science', 'the course');")
	database.DB.Exec("insert into courses (created_at, updated_at, name, description)values(datetime(),datetime(),'SE', 'course');")
	database.DB.Exec("insert into active_courses (instructor_id, course_id, start_date, end_date, created_at, updated_at)values(1, 1, datetime(), datetime(), datetime(), datetime());")
	database.DB.Exec("insert into active_courses (instructor_id, course_id, start_date, end_date, created_at, updated_at)values(1, 2, datetime(), datetime(), datetime(), datetime());")
	database.DB.Exec("insert into enrollments (user_id, active_course_id, status, enrollment_date, approval_date)values(2,1, 'approved', datetime(), datetime());")
	database.DB.Exec("insert into enrollments (user_id, active_course_id, status, enrollment_date, approval_date)values(2,2, 'approved', datetime(), datetime());")
	// database.DB.Exec("DELETE FROM users") // Clear users table
	// database.DB.Exec("DELETE FROM roles") // Clear roles table
}

func CloseTestDB() {
	database.DB.Exec("update sqlite_sequence set seq = 0")
	database.DB.Exec("DELETE FROM enrollments")
	database.DB.Exec("DELETE FROM user_roles")
	database.DB.Exec("DELETE FROM roles")
	database.DB.Exec("DELETE FROM active_courses")
	database.DB.Exec("DELETE FROM courses")
	database.DB.Exec("DELETE FROM users")
=======
	database.DB.AutoMigrate(&models.User{}) // Create schema
	database.DB.Exec("insert into roles (created_at, updated_At, name) values(datetime('now'),datetime('now'),'student');")
	database.DB.Exec("insert into roles (created_at, updated_At, name) values(datetime('now'),datetime('now'),'admin');")
	database.DB.Exec("insert into roles (created_at, updated_At, name) values(datetime('now'),datetime('now'),'instructor');")
	database.DB.Exec("insert into roles (created_at, updated_At, name) values(datetime('now'),datetime('now'),'teaching_assistant');")
	database.DB.Exec("DELETE FROM users") // Clear users table
	database.DB.Exec("DELETE FROM roles") // Clear roles table

	instructor := models.User{
		Username: "testinstructor",
		Email:    "testinstructor@example.com",
		Password: "password123",
	}
	if err := database.DB.Create(&instructor).Error; err != nil {
		fmt.Printf("Failed to create instructor user: %v\n", err)
	}

	for i := 1; i <= 30; i++ {

		course := models.Course{
			Name:        fmt.Sprintf("Course %d", i),
			Description: fmt.Sprintf("Description for course %d", i),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now().AddDate(0, 1, 0), // Ends in 1 month
		}
		database.DB.Create(&course)
		// Insert into Active Courses
		activeCourse := models.ActiveCourse{
			InstructorID: instructor.ID,
			CourseID:     course.ID,
			StartDate:    time.Now(),
			EndDate:      time.Now().AddDate(0, 1, 0), // Ends in 1 month
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			IsActive:     true,
			Capacity:     50 + i, // Random capacity values
			Enrolled:     0,
		}
		database.DB.Create(&activeCourse)
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
