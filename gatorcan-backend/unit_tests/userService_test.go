package unit

import (
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository mocks the user repository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsernameorEmail(username, email string) (*models.User, error) {
	args := m.Called(username, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) CreateNewUser(userData *dtos.UserCreateDTO) (*models.User, error) {
	args := m.Called(userData)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// MockRoleRepository mocks the role repository
type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) GetRolesByName(roles []string) ([]models.Role, error) {
	args := m.Called(roles)
	return args.Get(0).([]models.Role), args.Error(1)
}

var userRepo *MockUserRepository

func mockLogin(loginData *dtos.LoginRequestDTO) (*dtos.LoginResponseDTO, error) {
	user, err := userRepo.GetUserByUsername(loginData.Username)
	if err != nil {
		return &dtos.LoginResponseDTO{Err: true}, err
	}
	// Simulate password check and token generation
	if loginData.Password == "password123" && user != nil {
		return &dtos.LoginResponseDTO{Token: "mockToken", Err: false}, nil
	}
	return &dtos.LoginResponseDTO{Err: true}, errors.New("invalid credentials")
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	tests := []struct {
		name        string
		loginData   *dtos.LoginRequestDTO
		mockUser    *models.User
		mockError   error
		expectError bool
	}{
		{
			name: "Successful Login",
			loginData: &dtos.LoginRequestDTO{
				Username: "testuser",
				Password: "password123",
			},
			mockUser: &models.User{
				Username: "testuser",
				Password: "$2a$10$somehashedpassword", // Use a real hashed password here
				Roles:    []*models.Role{{Name: "user"}},
			},
			mockError:   nil,
			expectError: false,
		},
		{
			name: "User Not Found",
			loginData: &dtos.LoginRequestDTO{
				Username: "nonexistent",
			},
			mockUser:    nil,
			mockError:   errors.New("user not found"),
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.On("GetUserByUsername", tc.loginData.Username).Return(tc.mockUser, tc.mockError)

			response, err := mockLogin(tc.loginData)

			if tc.expectError {
				assert.Error(t, err)
				assert.True(t, response.Err)
			} else {
				assert.NoError(t, err)
				assert.False(t, response.Err)
				assert.NotEmpty(t, response.Token)
			}
		})
	}
}

func GetUserDetailsFromService(username string) (*models.User, error) {
	user, err := userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func TestGetUserDetails(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	mockUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		//CreatedAt: time.Now(),
		Roles: []*models.Role{{Name: "user"}},
	}

	tests := []struct {
		name        string
		username    string
		mockUser    *models.User
		mockError   error
		expectError bool
	}{
		{
			name:        "Success",
			username:    "testuser",
			mockUser:    mockUser,
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "User Not Found",
			username:    "nonexistent",
			mockUser:    nil,
			mockError:   errors.New("user not found"),
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.On("GetUserByUsername", tc.username).Return(tc.mockUser, tc.mockError)

			user, err := GetUserDetailsFromService(tc.username)

			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tc.mockUser.Username, user.Username)
			}
		})
	}
}

func UpdateUser(username string, updateData *dtos.UpdateUserDTO) error {
	user, err := userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	// Simulate password update
	if updateData.OldPassword == "oldpass" {
		user.Password = updateData.NewPassword
		return userRepo.UpdateUser(user)
	}
	return errors.New("incorrect old password")
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	tests := []struct {
		name        string
		username    string
		updateData  *dtos.UpdateUserDTO
		mockUser    *models.User
		mockError   error
		expectError bool
	}{
		{
			name:     "Success",
			username: "testuser",
			updateData: &dtos.UpdateUserDTO{
				OldPassword: "oldpass",
				NewPassword: "newpass",
			},
			mockUser: &models.User{
				Username: "testuser",
				Password: "$2a$10$somehashedpassword", // Use a real hashed password here
			},
			mockError:   nil,
			expectError: false,
		},
		{
			name:     "User Not Found",
			username: "nonexistent",
			updateData: &dtos.UpdateUserDTO{
				OldPassword: "oldpass",
				NewPassword: "newpass",
			},
			mockUser:    nil,
			mockError:   errors.New("user not found"),
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.On("GetUserByUsername", tc.username).Return(tc.mockUser, tc.mockError)
			if tc.mockUser != nil {
				mockRepo.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(nil)
			}

			err := UpdateUser(tc.username, tc.updateData)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
