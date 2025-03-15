package controllers

import (
	"context"
	"fmt"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/interfaces"
	"gatorcan-backend/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService interfaces.UserService
	logger      *log.Logger
}

func NewUserController(userService interfaces.UserService, logger *log.Logger) *UserController {
	return &UserController{
		userService: userService,
		logger:      logger,
	}
}

func (uc *UserController) CreateUser(c *gin.Context, logger *log.Logger) {

	logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

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

	response, err := uc.userService.CreateUser(ctx, userRequest)
	if err != nil {
		c.JSON(response.Code, gin.H{"error": response.Message})
		logger.Printf("Error in CreateUser service: %v", err)
		return
	}

	logger.Printf("User created successfully: %s", userRequest.Username)
	c.JSON(response.Code, gin.H{"message": response.Message})
}

func (uc *UserController) Login(c *gin.Context, logger *log.Logger) {

	logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var loginData *dtos.LoginRequestDTO

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		logger.Printf("Failed to bind JSON data: %v %d", err, c.Writer.Status())
		return
	}

	response, err := uc.userService.Login(ctx, loginData)
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

func (uc *UserController) GetUserDetails(c *gin.Context, logger *log.Logger) {
	logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	username := c.Param("username")
	user, err := uc.userService.GetUserDetails(ctx, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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

func (uc *UserController) DeleteUser(c *gin.Context, logger *log.Logger) {

	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	username := c.Param("username")

	logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	err := uc.userService.DeleteUser(ctx, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		logger.Printf("Error in Deleting User: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s has been deleted successfully", username)})
}

func (uc *UserController) UpdateUser(c *gin.Context, logger *log.Logger) {

	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

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

	err := uc.userService.UpdateUser(ctx, username, &updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Printf("Error in UpdateUser service: %v", err)
		return
	}

	logger.Printf("User updated successfully: %s", username)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User updated successfully: %s", username)})
}

func (uc *UserController) UpdateRoles(c *gin.Context, logger *log.Logger) {

	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var updateRolesDTO dtos.UpdateUserRolesDTO
	if err := c.ShouldBindJSON(&updateRolesDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := uc.userService.UpdateRoles(ctx, updateRolesDTO.Username, updateRolesDTO.Roles)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User roles updated successfully for %s", updateRolesDTO.Username)})
}

func UploadAssignments(c *gin.Context, logger *log.Logger) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "156 : " + err.Error()})
		return
	}
	fileHeader, err := utils.ValidateFile(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "161 : " + err.Error()})
		return
	}

	fmt.Println("dst : ", fileHeader.Path)

	userID, _ := c.Get("user_id")
	if err := services.UploadAssignments(fileHeader, userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "169 : " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user assesmet created with success."})
}
