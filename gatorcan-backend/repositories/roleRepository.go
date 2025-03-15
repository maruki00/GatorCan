package repositories

import (
	"context"
	"gatorcan-backend/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRolesByName(ctx context.Context, roleNames []string) ([]models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetRolesByName(ctx context.Context, roleNames []string) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.WithContext(ctx).
		Where("name IN ?", roleNames).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}
