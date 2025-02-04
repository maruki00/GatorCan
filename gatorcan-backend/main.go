package main

import (
	"fmt"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/routes"
	"gatorcan-backend/utils"

	"github.com/gin-gonic/gin"
)

func GenerateadminToken() {
	username := "muthu"
	role := []string{"admin", "TA"}

	token, err := utils.GenerateToken(username, role)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated JWT Token:", token)
}

func main() {

	database.Connect()

	GenerateadminToken()

	database.DB.AutoMigrate(&models.User{})

	// Set up router
	router := gin.Default()

	// Register routes
	routes.UserRoutes(router)

	router.Run(":8080")
}
