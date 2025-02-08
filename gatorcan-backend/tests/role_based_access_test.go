package tests

import (
	"bytes"
	"gatorcan-backend/models"
	"gatorcan-backend/routes"
	"gatorcan-backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.UserRoutes(router)
	return router
}

func generateToken(username string, roles []string) string {
	token, _ := utils.GenerateToken(username, roles)
	return token
}

func TestAdminAccess(t *testing.T) {
	router := setupRouter()

	adminToken := generateToken("adminUser", []string{string(models.Admin)})

	t.Run("CreateUser as Admin", func(t *testing.T) {
		w := httptest.NewRecorder()
		reqBody := bytes.NewBufferString(`{"username":"testuser","email":"testuser@example.com","password":"password123","roles":["student"]}`)
		req, _ := http.NewRequest("POST", "/user/add_user", reqBody)
		req.Header.Set("Authorization", "Bearer "+adminToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GetUserDetails as Admin", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/user/testuser", nil)
		req.Header.Set("Authorization", "Bearer "+adminToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("DeleteUser as Admin", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/user/testuser", nil)
		req.Header.Set("Authorization", "Bearer "+adminToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestInstructorAccess(t *testing.T) {
	router := setupRouter()

	instructorToken := generateToken("instructorUser", []string{string(models.Instructor)})

	t.Run("Access Instructor Route", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/instructor/upload-assignment", nil)
		req.Header.Set("Authorization", "Bearer "+instructorToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code) // Assuming the route is not implemented
	})

	t.Run("Access Admin Route as Instructor", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/add_user", nil)
		req.Header.Set("Authorization", "Bearer "+instructorToken)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}

func TestUnauthorizedAccess(t *testing.T) {
	router := setupRouter()

	t.Run("Access Admin Route without Token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user/add_user", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Access Instructor Route without Token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/instructor/upload-assignment", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
