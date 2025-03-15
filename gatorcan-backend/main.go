package main

import (
	"gatorcan-backend/config"
	"gatorcan-backend/controllers"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/repositories"
	"gatorcan-backend/routes"
	"gatorcan-backend/services"
	"gatorcan-backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {

	logger := utils.Log()

	logger.Println("Application started")

	// Load configuration
	appConfig := config.LoadConfig()
	logger.Printf("Environment: %s", appConfig.Environment)

	// Set Gin mode based on environment
	if appConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database
	db, err := database.Connect(appConfig.Database)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate all models
	logger.Println("Migrating database schema...")
	err = database.Migrate(db,
		&models.User{},
		&models.Role{},
		&models.Course{},
		&models.ActiveCourse{},
		&models.Enrollment{},
	)
	if err != nil {
		logger.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	courseRepo := repositories.NewCourseRepository(db)
	roleRepo := repositories.NewRoleRepository(db) // Add missing role repository

	// Initialize HTTP client with sensible defaults
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	// Initialize services with consistent pattern
	userService := services.NewUserService(courseRepo, userRepo, roleRepo, appConfig, httpClient)
	courseService := services.NewCourseService(courseRepo, userRepo, appConfig, httpClient)

	// Initialize controllers
	userController := controllers.NewUserController(userService, logger)
	courseController := controllers.NewCourseController(courseService, logger)

	// Set up router
	router := gin.Default()

	// Register routes
	routes.UserRoutes(
		userController,
		courseController,
		router,
		logger)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	http.ListenAndServe(":8080", handler)

	//router.Run(":8080")
}
