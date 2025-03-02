package controllers

import (
	"fmt"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/services"
	"gatorcan-backend/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context, logger *log.Logger) {

	logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	var userRequest *dtos.UserRequestDTO

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		logger.Printf("Failed to bind JSON data: %v %d", err, c.Writer.Status())
		return
	}

	if !utils.IsValidEmail(userRequest.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	response, err := services.CreateUser(userRequest)
	if err != nil {
		c.JSON(response.Code, gin.H{"error": response.Message})
		logger.Printf("Error in CreateUser service: %v", err)
		return
	}

	logger.Printf("User created successfully: %s", userRequest.Username)
	c.JSON(response.Code, gin.H{"message": response.Message})
}

func Login(c *gin.Context, logger *log.Logger) {

	logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	var loginData *dtos.LoginRequestDTO

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.Printf("Failed to bind JSON data: %v %d", err, c.Writer.Status())
		return
	}

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

func GetUserDetails(c *gin.Context, logger *log.Logger) {
	logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	username := c.Param("username")
	user, err := services.GetUserDetails(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Printf("Error in GetUserDetails service: %v", err)
		return
	}
	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"username":   user.Username,
		"email":      user.Email,
		"roles":      roleNames,
		"created_at": user.CreatedAt,
	})
}

func DeleteUser(c *gin.Context, logger *log.Logger) {

	username := c.Param("username")

	logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	err := services.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Printf("Error in Deleting User: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s has been deleted successfully", username)})
}

func UpdateUser(c *gin.Context, logger *log.Logger) {
	var updateData dtos.UpdateUserDTO
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		logger.Printf("Failed to bind JSON data: %v", err)
		return
	}

	usernameInterface, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	username, ok := usernameInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username in context"})
		return
	}

	err := services.UpdateUser(username, &updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Printf("Error in UpdateUser service: %v", err)
		return
	}

	logger.Printf("User updated successfully: %s", username)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User updated successfully: %s", username)})
}

func UpdateRoles(c *gin.Context, logger *log.Logger) {
	var updateRolesDTO dtos.UpdateUserRolesDTO
	if err := c.ShouldBindJSON(&updateRolesDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := services.UpdateRoles(updateRolesDTO.Username, updateRolesDTO.Roles)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User roles updated successfully for %s", updateRolesDTO.Username)})
}
