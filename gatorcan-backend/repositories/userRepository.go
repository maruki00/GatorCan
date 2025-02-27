package repositories

import (
	"gatorcan-backend/database"
	"gatorcan-backend/models"
)

type UserRepository interface {
	// Interface Method Declarations
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type userRepository struct {
	// todo: database connection or ORM instance

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
