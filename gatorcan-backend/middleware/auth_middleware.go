package middleware

import (
	"fmt"
	"gatorcan-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token
func AuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			fmt.Println("No Authorization token received")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		fmt.Println("🔍 Received Token 1:", tokenString)

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			fmt.Println("Invalid Token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		for _, role := range claims.Roles {
			for _, requiredRole := range requiredRoles {
				if role == requiredRole {
					fmt.Println(role + "==" + requiredRole)
					c.Set("username", claims.Username)
					c.Set("roles", claims.Roles)
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
	}
}
