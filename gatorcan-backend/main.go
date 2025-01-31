package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to GatorCan Backend!"})
	})

	fmt.Println("Server is running on port 8080")
	r.Run(":8080") // Start the server on port 8080
}
