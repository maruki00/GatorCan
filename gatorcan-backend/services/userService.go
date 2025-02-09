package services

import (
	"errors"
	"gatorcan-backend/models"
	"gatorcan-backend/repositories"
	"gatorcan-backend/utils"
	"net/http"
)

type loginResponse struct {
	Err     bool
	Code    int
	Message string
	Token   string
}

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(loginData *LoginData) (*loginResponse, error) {
	var response loginResponse
	var user *models.User

	// get user from the database
	user, err := repositories.NewUserRepository().GetUserByUsername(loginData.Username)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Failed to get user data"
		response.Err = true
		return &response, err
	}

	// Check if the user exists
	if user == nil {
		response.Code = http.StatusUnauthorized
		response.Message = "Invalid username or password"
		response.Err = true
		return &response, err
	}

	// Check if the password matches
	if err := utils.VerifyPassword(user.Password, loginData.Password); !err {
		response.Code = http.StatusUnauthorized
		response.Message = "Invalid username or password"
		response.Err = true
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
		response.Code = http.StatusInternalServerError
		response.Message = "Failed to generate token"
		response.Err = true
		return &response, err
	}

	// The token should be set as authorization header
	response.Code = http.StatusOK
	response.Message = "Login successful"
	response.Err = false
	response.Token = token
	return &response, nil
}
