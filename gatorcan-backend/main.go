package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/routes"
	"gatorcan-backend/utils"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "gen-token" {
		utils.GenerateadminToken()
		return
	}

	database.Connect()

	database.DB.AutoMigrate(&models.User{})

	// Set up router
	router := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(router)

	router.Run(":8080")
}
