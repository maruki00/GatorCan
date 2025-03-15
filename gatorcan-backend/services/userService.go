package services

import (
	"context"
	"errors"
	"fmt"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	"gatorcan-backend/interfaces"
	"gatorcan-backend/models"
	"gatorcan-backend/utils"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	courseRepo interfaces.CourseRepository
	userRepo   interfaces.UserRepository
	roleRepo   interfaces.RoleRepository
	config     *config.AppConfig
	httpClient interfaces.HTTPClient
}

func NewUserService(
	courseRepo interfaces.CourseRepository,
	userRepo interfaces.UserRepository,
	roleRepo interfaces.RoleRepository,
	config *config.AppConfig,
	httpClient interfaces.HTTPClient,
) interfaces.UserService {
	return &UserServiceImpl{
		courseRepo: courseRepo,
		userRepo:   userRepo,
		roleRepo:   roleRepo,
		config:     config,
		httpClient: httpClient,
	}
}

func (s *UserServiceImpl) Login(ctx context.Context, loginData *dtos.LoginRequestDTO) (*dtos.LoginResponseDTO, error) {
	var response dtos.LoginResponseDTO
	var user *models.User

	// get user from the database
	user, err := s.userRepo.GetUserByUsername(ctx, loginData.Username)
	if err != nil {
		response.Code = http.StatusNotFound
		response.Message = "Invalid_username"
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

func (s *UserServiceImpl) CreateUser(ctx context.Context, userData *dtos.UserRequestDTO) (*dtos.UserResponseDTO, error) {
	var response dtos.UserResponseDTO

	//roleRepo := repositories.NewRolesRepository()

	existingUser, err := s.userRepo.GetUserByUsernameorEmail(ctx, userData.Username, userData.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			existingUser = nil
		} else {
			response.Code = http.StatusInternalServerError
			response.Message = "Failed to check user existence"
			response.Err = true
			return &response, err
		}
	}
	if existingUser != nil {
		response.Code = http.StatusBadRequest
		response.Message = "User already exists"
		response.Err = true
		return &response, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Failed to hash password"
		response.Err = true
		return &response, err
	}

	// Fetch roles from database
	newUserRoles, err := s.roleRepo.GetRolesByName(ctx, userData.Roles)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "One or more roles not found"
		response.Err = true
		return &response, err
	}

	var newUserRolesPtrs []*models.Role
	for _, role := range newUserRoles {
		newUserRolesPtrs = append(newUserRolesPtrs, &role)
	}

	// Prepare DTO for user creation
	userCreateDTO := &dtos.UserCreateDTO{
		Username: userData.Username,
		Email:    userData.Email,
		Password: hashedPassword,
		Roles:    newUserRolesPtrs, // Already a slice of `models.Role`
	}

	// Create new user
	_, err = s.userRepo.CreateNewUser(ctx, userCreateDTO)
	// Create new user
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Failed to create user"
		response.Err = true
		return &response, err
	}

	response.Code = http.StatusCreated
	response.Message = "User created successfully"
	response.Err = false
	return &response, nil
}

func (s *UserServiceImpl) GetUserDetails(ctx context.Context, username string) (*models.User, error) {

	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, username string) error {

	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}
	err = s.userRepo.DeleteUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, username string, updateData *dtos.UpdateUserDTO) error {

	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	if !utils.VerifyPassword(user.Password, updateData.OldPassword) {
		return errors.New("incorrect old password")
	}

	hashedPassword, err := utils.HashPassword(updateData.NewPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) UpdateRoles(ctx context.Context, username string, roles []string) error {

	// roleRepo := repositories.NewRolesRepository()

	// Fetch user
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	// Fetch roles in a single query
	newRoles, err := s.roleRepo.GetRolesByName(ctx, roles)
	if err != nil {
		return err
	}

	// Check for missing roles
	foundRoles := make(map[string]bool)
	for _, role := range newRoles {
		foundRoles[role.Name] = true
	}

	var missingRoles []string
	for _, role := range roles {
		if !foundRoles[role] {
			missingRoles = append(missingRoles, role)
		}
	}

	if len(missingRoles) > 0 {
		return fmt.Errorf("roles not found: %v", missingRoles)
	}

	// Update user's roles
	user.Roles = make([]*models.Role, len(newRoles))
	for i := range newRoles {
		user.Roles[i] = &newRoles[i]
	}

	err = s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func UploadAssignments(fileHeader *utils.FileHeader, user_id uint) error {

	S3, err := utils.NewS3()
	if err != nil {
		return fmt.Errorf("could not create s3 object")
	}
	info, _ := os.Stat(fileHeader.Path)

	fmt.Println("dst service : ", fileHeader.Path)
	err = S3.UploadFile(context.TODO(), info.Name(), fileHeader.Path, fileHeader.ContentType, true)
	if err != nil {
		return fmt.Errorf("could not upload the file, %s", err.Error())
	}
	userRepo := repositories.NewUserRepository()

	if err := userRepo.CreateAssignment(fileHeader.Path, user_id); err != nil {
		return fmt.Errorf("could not create file meta data")
	}

	return nil
}
