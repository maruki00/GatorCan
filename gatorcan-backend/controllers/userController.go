package controllers

import (
	"errors"
	"fmt"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"gatorcan-backend/repositories"
	"gatorcan-backend/services"
	"gatorcan-backend/utils"
	"log"
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

	var user UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if !utils.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	var count int64
	database.DB.Model(&models.User{}).Where("username = ? OR email = ?", user.Username, user.Email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	var newUserRoles []*models.Role
	if err := database.DB.Where("name IN ?", user.Roles).Find(&newUserRoles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "One or more roles not found"})
		return
	}

	newUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
		Roles:    newUserRoles,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	log.Printf("User created: %s, Email: %s", newUser.Username, newUser.Email)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s has been created successfully", newUser.Username)})
}

// Handler function for the login route
func Login(c *gin.Context, logger *log.Logger) {

	logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	var loginData *dtos.LoginRequestDTO

	// Bind JSON data to loginData struct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.Printf("Failed to bind JSON data: %v %d", err, c.Writer.Status())
		return
	}

	// Call the login service
	response, err := services.Login(loginData)
	if response.Err || err != nil {
		c.JSON(response.Code, gin.H{"error": response.Message})
		logger.Printf("Login Service Error: %v %d", err, c.Writer.Status())
	} else {
		c.Writer.Header().Set("Authorization", "Bearer "+response.Token)
		c.JSON(response.Code, gin.H{
			"message": response.Message,
			"token":   response.Token,
		})
		logger.Printf("Response: %s %s %d", c.Request.Method, c.Request.URL.Path, c.Writer.Status())
	}
}

func GetUserDetails(c *gin.Context) {
	// Get the username from the route parameter
	username := c.Param("username")

	// Query the database to get the user by username, including their roles
	var user models.User
	if err := database.DB.Preload("Roles").Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Create a slice to hold role names
	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}

	// Return user details including the roles (excluding password for security reasons)
	c.JSON(http.StatusOK, gin.H{
		"username":   user.Username,
		"email":      user.Email,
		"roles":      roleNames, // Use the slice of role names
		"created_at": user.CreatedAt,
	})
}

func DeleteUser(c *gin.Context) {

	username := c.Param("username")

	// Check if the user exists before deleting
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s has been deleted successfully", username)})
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
	fmt.Println(username, "aaaaaa")
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

	// Validate request body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Fetch user
	var user models.User
	if err := database.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Fetch roles in a single query
	var roles []models.Role
	if err := database.DB.Where("name IN (?)", request.Roles).Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch roles"})
		return
	}

	// Check for missing roles
	foundRoles := make(map[string]bool)
	for _, role := range roles {
		foundRoles[role.Name] = true
	}

	var missingRoles []string
	for _, role := range request.Roles {
		if !foundRoles[role] {
			missingRoles = append(missingRoles, role)
		}
	}

	if len(missingRoles) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Roles not found: %v", missingRoles)})
		return
	}

	// Update user's roles
	user.Roles = make([]*models.Role, len(roles))
	for i := range roles {
		user.Roles[i] = &roles[i]
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User roles updated successfully"})
}

func GetEnrolledCourses(c *gin.Context, logger *log.Logger) {
	// Get username from JWT token
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	enrollments, err := repositories.NewCourseRepository().GetEnrolledCourses(username.(string))
	if err == errors.New("user not found") {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		logger.Printf("user not found: %s %d", username, c.Writer.Status())
		return
	} else if err == errors.New("failed to fetch enrolled courses") {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrolled courses"})
		logger.Printf("failed to fetch enrolled courses: %s %d", username, c.Writer.Status())
		return
	}

	// Return enrolled courses
	c.JSON(http.StatusOK, enrollments)
}
