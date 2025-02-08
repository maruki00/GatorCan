package middleware

import (
	"gatorcan-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token and checks for required roles
func AuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Store token claims in context
		c.Set("username", claims.Username)
		c.Set("roles", claims.Roles)

		// If no roles are required, allow access
		if len(requiredRoles) == 0 {
			c.Next()
			return
		}

		// Convert required roles into a set for efficient lookup
		requiredRolesMap := make(map[string]struct{}, len(requiredRoles))
		for _, role := range requiredRoles {
			requiredRolesMap[role] = struct{}{}
		}

		// Check if the user has any required role
		for _, userRole := range claims.Roles {
			if _, exists := requiredRolesMap[userRole]; exists {
				c.Next()
				return
			}
		}

		// If no matching role is found, deny access
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
	}
}
