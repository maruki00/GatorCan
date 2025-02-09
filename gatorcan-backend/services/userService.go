package services

import (
	"errors"
	"gatorcan-backend/models"
	"gatorcan-backend/repositories"
	"gatorcan-backend/utils"
	"net/http"
)

type loginResponse struct {
	error   bool
	code    int
	message string
	token   string
}

func Login() (*loginResponse, error) {
	var response loginResponse
	var user *models.User
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// get user from the database
	user, err := repositories.NewUserRepository().GetUserByUsername(loginData.Username)
	if err != nil {
		response.code = http.StatusInternalServerError
		response.message = "Failed to get user data"
		response.error = true
		return &response, err
	}

	// Check if the user exists
	if user == nil {
		response.code = http.StatusUnauthorized
		response.message = "Invalid username or password"
		response.error = true
		return &response, err
	}

	// Check if the password matches
	if err := utils.VerifyPassword(user.Password, loginData.Password); !err {
		response.code = http.StatusUnauthorized
		response.message = "Invalid username or password"
		response.error = true
		return &response, errors.New("Invalid username or password")
	}

	// Convert roles to []string
	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}

	// Generate JWT token
	token, err := utils.GenerateToken(loginData.Username, roleNames)
	if err != nil {
		response.code = http.StatusInternalServerError
		response.message = "Failed to generate token"
		response.error = true
		return &response, err
	}

	// The token should be set as authorization header
	response.code = http.StatusOK
	response.message = "Login successful"
	response.error = false
	response.token = token
	return &response, nil
}
