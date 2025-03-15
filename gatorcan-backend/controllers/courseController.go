package controllers

import (
	"context"
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/interfaces"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	courseService interfaces.CourseService
	logger        *log.Logger
}

func NewCourseController(service interfaces.CourseService, logger *log.Logger) *CourseController {
	return &CourseController{
		courseService: service,
		logger:        logger,
	}
}

//	func (cc *CourseController) GetCoursesService() *services.CourseService {
//		Repo := repositories.NewCourseRepository()
//		courseService := services.NewCourseService(Repo)
//		return courseService
//	}
func (cc *CourseController) GetEnrolledCourses(c *gin.Context) {
	cc.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	//courseService := GetCoursesService()

	// Get username from JWT token
	username, exists := c.Get("username")
	if !exists {
		cc.logger.Printf("Unauthorized access attempt to enrolled courses")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	enrollments, err := cc.courseService.GetEnrolledCourses(ctx, cc.logger, username.(string))
	if err == errors.New("user not found") {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrolled courses"})
		return
	}

	// Return enrolled courses
	c.JSON(http.StatusOK, enrollments)
}

func (cc *CourseController) GetCourses(c *gin.Context) {
	cc.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	//courseService := GetCoursesService()

	username, exists := c.Get("username")
	if !exists {
		cc.logger.Printf("Unauthorized access attempt to courses")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// Get query parameters with default values
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Call the service layer to fetch courses
	courses, err := cc.courseService.GetCourses(ctx, cc.logger, username.(string), page, pageSize)
	if err != nil {
		cc.logger.Printf("Failed to fetch courses for page %d with pageSize %d: %v", page, pageSize, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	// Return courses
	c.JSON(http.StatusOK, courses)
}

func (cc *CourseController) EnrollInCourse(c *gin.Context) {
	cc.logger.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var request dtos.EnrollRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate courseID
	courseID := request.CourseID
	if courseID < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	err := cc.courseService.EnrollUser(ctx, cc.logger, username.(string), courseID)
	if err == errors.New("user not found") {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err == errors.New("course not found") {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	} else if err == errors.New("enrollment request already exists") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enrollment request already exists"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to request enrollment"})
		return
	}

	// Respond with a success message using the EnrollmentResponseDTO
	response := dtos.EnrollmentResponseDTO{
		Message: "Enrollment requested successfully",
	}

	c.JSON(http.StatusCreated, response)
}
