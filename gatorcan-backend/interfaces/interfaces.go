package interfaces

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/models"
	"log"
	"net/http"
)

type CourseRepository interface {
	GetEnrolledCourses(ctx context.Context, userID int) ([]models.Enrollment, error)
	GetCourses(ctx context.Context, page, pageSize int) ([]models.Course, error)
	GetCourseByID(ctx context.Context, courseID int) (models.ActiveCourse, error)
	RequestEnrollment(ctx context.Context, userID, activeCourseID uint) error
	ApproveEnrollment(ctx context.Context, enrollmentID uint) error
	RejectEnrollment(ctx context.Context, enrollmentID uint) error
	GetPendingEnrollments(ctx context.Context) ([]models.Enrollment, error)
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByUsernameorEmail(ctx context.Context, username string, email string) (*models.User, error)
	CreateNewUser(ctx context.Context, userDTO *dtos.UserCreateDTO) (*models.User, error)
	DeleteUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	UpdateUserRoles(ctx context.Context, user *models.User, roles []*models.Role) error
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type CourseService interface {
	GetEnrolledCourses(ctx context.Context, logger *log.Logger, username string) ([]dtos.EnrolledCoursesResponseDTO, error)
	GetCourses(ctx context.Context, logger *log.Logger, username string, page, pageSize int) ([]dtos.CourseResponseDTO, error)
	EnrollUser(ctx context.Context, logger *log.Logger, username string, courseID int) error
}

type UserService interface {
	Login(ctx context.Context, loginData *dtos.LoginRequestDTO) (*dtos.LoginResponseDTO, error)
	GetUserDetails(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, username string, user *dtos.UpdateUserDTO) error
	CreateUser(ctx context.Context, user *dtos.UserRequestDTO) (*dtos.UserResponseDTO, error)
	DeleteUser(ctx context.Context, username string) error
	UpdateRoles(ctx context.Context, username string, roles []string) error
}

type RoleRepository interface {
	GetRolesByName(ctx context.Context, roleNames []string) ([]models.Role, error)
}
