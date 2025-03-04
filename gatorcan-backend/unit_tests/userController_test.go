package unit

import (
	"bytes"
	"encoding/json"
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService mocks the user service
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(dto *dtos.UserRequestDTO) (*dtos.ResponseDTO, error) {
	args := m.Called(dto)
	return args.Get(0).(*dtos.ResponseDTO), args.Error(1)
}

func (m *MockUserService) Login(dto *dtos.LoginRequestDTO) (*dtos.LoginResponseDTO, error) {
	args := m.Called(dto)
	return args.Get(0).(*dtos.LoginResponseDTO), args.Error(1)
}

func (m *MockUserService) GetUserDetails(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func Login(c *gin.Context, logger *log.Logger) {
	var loginRequest dtos.LoginRequestDTO
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mockService := new(MockUserService)
	response, err := mockService.Login(&loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(response.Code, gin.H{"message": response.Message, "token": response.Token})
}

func CreateUser(c *gin.Context, logger *log.Logger) {
	var userRequest dtos.UserRequestDTO
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mockService := new(MockUserService)
	response, err := mockService.CreateUser(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(response.Code, gin.H{"message": response.Message})
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(bytes.NewBuffer([]byte{}), "", log.LstdFlags)

	tests := []struct {
		name           string
		input          dtos.UserRequestDTO
		mockResponse   *dtos.ResponseDTO
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			input: dtos.UserRequestDTO{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			mockResponse: &dtos.ResponseDTO{
				Code:    http.StatusCreated,
				Message: "User created successfully",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"message":"User created successfully"}`,
		},
		{
			name: "Invalid Email",
			input: dtos.UserRequestDTO{
				Username: "testuser",
				Email:    "invalid-email",
				Password: "password123",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid email format"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockUserService)
			if tc.mockResponse != nil {
				mockService.On("CreateUser", &tc.input).Return(tc.mockResponse, tc.mockError)
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonInput, _ := json.Marshal(tc.input)
			c.Request = httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonInput))
			c.Request.Header.Set("Content-Type", "application/json")

			CreateUser(c, logger)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())
		})
	}
}

func TestUserControllerLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(bytes.NewBuffer([]byte{}), "", log.LstdFlags)

	tests := []struct {
		name           string
		input          dtos.LoginRequestDTO
		mockResponse   *dtos.LoginResponseDTO
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			input: dtos.LoginRequestDTO{
				Username: "testuser",
				Password: "password123",
			},
			mockResponse: &dtos.LoginResponseDTO{
				Code:    http.StatusOK,
				Message: "Login successful",
				Token:   "jwt-token",
				Err:     false,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Login successful","token":"jwt-token"}`,
		},
		{
			name: "Invalid Credentials",
			input: dtos.LoginRequestDTO{
				Username: "testuser",
				Password: "wrongpassword",
			},
			mockResponse: &dtos.LoginResponseDTO{
				Code:    http.StatusUnauthorized,
				Message: "Invalid credentials",
				Err:     true,
			},
			mockError:      errors.New("invalid credentials"),
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid credentials"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockUserService)
			mockService.On("Login", &tc.input).Return(tc.mockResponse, tc.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonInput, _ := json.Marshal(tc.input)
			c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
			c.Request.Header.Set("Content-Type", "application/json")

			Login(c, logger)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())
		})
	}
}

func GetUserDetails(c *gin.Context, logger *log.Logger) {
	username := c.Param("username")

	mockService := new(MockUserService)
	user, err := mockService.GetUserDetails(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username, "email": user.Email, "roles": user.Roles})
}

func TestUserControllerGetUserDetails(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(bytes.NewBuffer([]byte{}), "", log.LstdFlags)

	mockUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		// CreatedAt: time.Now(),
		Roles: []*models.Role{{Name: "user"}},
	}

	tests := []struct {
		name           string
		username       string
		mockUser       *models.User
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success",
			username:       "testuser",
			mockUser:       mockUser,
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "User Not Found",
			username:       "nonexistent",
			mockUser:       nil,
			mockError:      errors.New("user not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"user not found"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockUserService)
			mockService.On("GetUserDetails", tc.username).Return(tc.mockUser, tc.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{{Key: "username", Value: tc.username}}

			GetUserDetails(c, logger)

			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			}
		})
	}
}
