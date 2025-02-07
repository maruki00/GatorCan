package repositories

import (
	"gatorcan-backend/database"
	"gatorcan-backend/models"
)

type UserRepository interface {
	// Interface Method Declarations
	GetUserByUsername(username string) (*models.User, error)
}

type userRepository struct {
	// todo: database connection or ORM instance

}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }
	// var user1 models.User
	// user = models.User{
	// 	Username: "admin",
	// 	Email:    "admin@email.com",
	// 	Password: string(hashedPassword),
	// 	Roles:    []string{"admin", "Instructor"},
	// }
	err := database.DB.Preload("Roles").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
