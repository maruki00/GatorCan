package repositories

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByUsernameorEmail(ctx context.Context, username string, email string) (*models.User, error)
	CreateNewUser(ctx context.Context, userDTO *dtos.UserCreateDTO) (*models.User, error)
	DeleteUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	UpdateUserRoles(ctx context.Context, user *models.User, roles []*models.Role) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Preload("Roles").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Preload("Roles").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsernameorEmail(ctx context.Context, username string, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Preload("Roles").Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepository) CreateNewUser(ctx context.Context, userDTO *dtos.UserCreateDTO) (*models.User, error) {
	newUser := models.User{
		Username: userDTO.Username,
		Email:    userDTO.Email,
		Password: userDTO.Password,
		Roles:    userDTO.Roles,
	}

	if err := r.db.WithContext(ctx).Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUserRoles(ctx context.Context, user *models.User, roles []*models.Role) error {
	err := r.db.WithContext(ctx).Model(&user).Association("Roles").Replace(roles)
	if err != nil {
		return err
	}
	return nil
}
