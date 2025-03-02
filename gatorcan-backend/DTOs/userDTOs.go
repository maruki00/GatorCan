package dtos

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
