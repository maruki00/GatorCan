package repositories

import (
	"gatorcan-backend/database"
	"gatorcan-backend/models"
)

type RoleRepository interface {
	GetRolesByName(roleNames []string) ([]models.Role, error)
}

type roleRepository struct {
}

func NewRolesRepository() RoleRepository {
	return &roleRepository{}
}

func (r *roleRepository) GetRolesByName(roleNames []string) ([]models.Role, error) {
	var roles []models.Role
	err := database.DB.Where("name IN ?", roleNames).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}
