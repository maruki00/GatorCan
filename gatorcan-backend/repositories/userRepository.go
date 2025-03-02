package repositories

import (
	"gatorcan-backend/database"
	"gatorcan-backend/models"
)

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsernameorEmail(username string, email string) (*models.User, error)
	CreateNewUser(username string, email string, password string, roles []*models.Role) (*models.User, error)
	DeleteUser(user *models.User) error
	UpdateUser(user *models.User) error
	UpdateUserRoles(user *models.User, roles []*models.Role) error
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Preload("Roles").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.Preload("Roles").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsernameorEmail(username string, email string) (*models.User, error) {
	var user models.User
	err := database.DB.Preload("Roles").Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepository) CreateNewUser(username string, email string, password string, roles []*models.Role) (*models.User, error) {
	var user models.User
	newUser := models.User{
		Username: username,
		Email:    email,
		Password: password,
		Roles:    roles,
	}
	err := database.DB.Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) DeleteUser(user *models.User) error {
	err := database.DB.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	err := database.DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUserRoles(user *models.User, roles []*models.Role) error {
	err := database.DB.Model(&user).Association("Roles").Replace(roles)
	if err != nil {
		return err
	}
	return nil
}
