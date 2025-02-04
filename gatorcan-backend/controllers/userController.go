package controllers

import (
	"fmt"
	"gatorcan-backend/database"
	"gatorcan-backend/middleware"
	"gatorcan-backend/models"
	"gatorcan-backend/repositories"
	"gatorcan-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username string   `json:"username" form:"username"`
	Email    string   `json:"email" form:"email"`
	Password string   `json:"password" form:"password"`
	Roles    []string `json:"roles" form:"roles"`
}

func CreateUser(c *gin.Context) {
	middleware.AuthMiddleware()(c)
	if c.IsAborted() {
		return
	}

	roles, exists := c.Get("roles")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: No roles found"})
		return
	}

	rolesSlice, ok := roles.([]string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid roles format"})
		return
	}

	if !hasRole(rolesSlice, "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Only admins can register users"})
		return
	}

	var user UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" || user.Roles == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing username, email, password or role"})
		return
	}

	var existingUser models.User
	if err := database.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	if !utils.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
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
		Roles:    user.Roles,
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

// Handler function for the login route
func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON data to loginData struct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user from the database
	user, err := repositories.NewUserRepository().GetUserByUsername(loginData.Username)

	// Check if the user exists
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	// Check if the password matches
	if err := utils.VerifyPassword(user.Password, loginData.Password); !err {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(loginData.Username, user.Roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// The token should be set as authorization header
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	// Return the token
	c.JSON(http.StatusOK, gin.H{"token": "generated"})

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}

func hasRole(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
