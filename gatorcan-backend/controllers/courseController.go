package controllers

import (
	"errors"
	"gatorcan-backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
