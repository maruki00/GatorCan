package dtos

type UpdateUserRolesDTO struct {
	Username string   `json:"username" binding:"required"`
	Roles    []string `json:"roles" binding:"required"`
}
