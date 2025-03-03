package controllers

import (
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEnrolledCourses(c *gin.Context, logger *log.Logger) {
	// Get username from JWT token
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	enrollments, err := services.GetEnrolledCourses(logger, username.(string))
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

func GetCourses(c *gin.Context, logger *log.Logger) {

	username, exists := c.Get("username")
	if !exists {
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
	courses, err := services.GetCourses(logger, username.(string), page, pageSize)
	if err != nil {
		logger.Printf("Failed to fetch courses for page %d with pageSize %d: %v", page, pageSize, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	// Return courses
	c.JSON(http.StatusOK, courses)
}

func EnrollInCourse(c *gin.Context, logger *log.Logger) {
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

	err := services.EnrollUser(logger, username.(string), courseID)
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
