package controllers

import (
	"fmt"
	"gatorcan-backend/database"
	"gatorcan-backend/middleware"
	"gatorcan-backend/models"
	"gatorcan-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func CreateUser(c *gin.Context) {
	middleware.AuthMiddleware()(c)
	if c.IsAborted() {
		return
	}

	role, exists := c.Get("role")
	if !exists || role.(string) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Only admins can register users"})
		return
	}

	var user UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing username, email, password or role"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
		Role:     user.Role,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	fmt.Printf("User %s has been created with email %s\n", newUser.Username, newUser.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User %s has been created successfully", newUser.Username),
	})
}
