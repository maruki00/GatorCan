package main

import (
	"github.com/gin-gonic/gin"

	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/routes"
)

func main() {

	database.Connect()

	database.DB.AutoMigrate(&models.User{})

	// Set up router
	router := gin.Default()

	// Register routes
	routes.UserRoutes(router)

	router.Run(":8080")
}
