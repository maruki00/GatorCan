package dtos

import "gatorcan-backend/models"

type UserRequestDTO struct {
	Username string   `json:"username" form:"username"`
	Email    string   `json:"email" form:"email"`
	Password string   `json:"password" form:"password"`
	Roles    []string `json:"roles" form:"roles"`
}

type UserResponseDTO struct {
	Err     bool
	Code    int
	Message string
}

type UserCreateDTO struct {
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Roles    []*models.Role `json:"roles"`
}
