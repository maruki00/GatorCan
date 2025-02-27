package controllers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetEnrolledCourses(t *testing.T) {
	logger := log.New(bytes.NewBuffer([]byte{}), "", log.LstdFlags)

	tests := []struct {
		name           string
		setupContext   func() *gin.Context
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Unauthorized",
			setupContext: func() *gin.Context {
				c, _ := gin.CreateTestContext(httptest.NewRecorder())
				return c
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Unauthorized"}`,
		},
		{
			name: "User not found",
			setupContext: func() *gin.Context {
				c, _ := gin.CreateTestContext(httptest.NewRecorder())
				c.Set("username", "nonexistentuser")
				return c
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"User not found"}`,
		},
		{
			name: "Internal server error",
			setupContext: func() *gin.Context {
				c, _ := gin.CreateTestContext(httptest.NewRecorder())
				c.Set("username", "internalerroruser")
				return c
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Failed to fetch enrolled courses"}`,
		},
		{
			name: "Success",
			setupContext: func() *gin.Context {
				c, _ := gin.CreateTestContext(httptest.NewRecorder())
				c.Set("username", "validuser")
				return c
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[]`, // Assuming the user has no enrolled courses for simplicity
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.setupContext()
			w := httptest.NewRecorder()
			c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

			GetEnrolledCourses(c, logger)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
