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

	if !utils.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	var existingUser models.User
	if err := database.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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

func GetUserDetails(c *gin.Context) {
	// Get the username from the route parameter
	username := c.Param("username")

	// Query the database to get the user by username
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return user details excluding password (for security reasons)
	c.JSON(http.StatusOK, gin.H{
		"username":   user.Username,
		"email":      user.Email,
		"roles":      user.Roles,
		"created_at": user.CreatedAt,
	})
}

func DeleteUser(c *gin.Context) {
	// Extract the username from the URL parameter
	username := c.Param("username")

	// Check for admin roles (optional, if only admin should delete users)
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
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Only admins can delete users"})
		return
	}

	// Query the database to find the user by username
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		// If no user is found, return a "user not found" message
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user from the database
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User %s has been deleted successfully", username),
	})
}

func UpdateUser(c *gin.Context) {
	var request struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get username from JWT token
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Fetch user from DB
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Verify old password
	if !utils.VerifyPassword(user.Password, request.OldPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect old password"})
		return
	}

	// Hash the new password
	hashedPassword, _ := utils.HashPassword(request.NewPassword)
	user.Password = hashedPassword

	// Save updated password
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func UpdateRoles(c *gin.Context) {
	var request struct {
		Username string   `json:"username" binding:"required"`
		Roles    []string `json:"roles" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the requester is an admin
	roles, _ := c.Get("roles") // Get user roles from JWT
	userRoles := roles.([]string)
	isAdmin := false
	for _, role := range userRoles {
		if role == "admin" {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can update roles"})
		return
	}

	// Fetch user from DB
	var user models.User
	if err := database.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update roles
	user.Roles = request.Roles
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User roles updated successfully"})
}
