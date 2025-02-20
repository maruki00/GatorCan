package main

import (
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/routes"
	"gatorcan-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	logger := utils.Log()

	logger.Println("Application started")

	database.Connect()

	database.DB.AutoMigrate(&models.User{})

	// Set up router
	router := gin.Default()

	// Register routes
	routes.UserRoutes(router, logger)

	router.Run(":8080")
}
