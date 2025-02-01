package main

import (
	"github.com/gin-gonic/gin"

	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/routes"
)

func main() {

	//uncomment this block to generate a token for the admin user
	// if len(os.Args) > 1 && os.Args[1] == "gen-token" {
	// 	utils.GenerateadminToken()
	// 	return
	// }

	database.Connect()

	database.DB.AutoMigrate(&models.User{})

	// Set up router
	router := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(router)

	router.Run(":8080")
}
