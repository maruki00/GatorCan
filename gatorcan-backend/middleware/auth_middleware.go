package middleware

import (
	"fmt"
	"gatorcan-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			fmt.Println("No Authorization token received")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		// 	tokenString = tokenString[7:]
		// }

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		fmt.Println("🔍 Received Token:", tokenString)

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			fmt.Println("Invalid Token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		fmt.Println("Decoded Token Claims:", claims)

		c.Set("username", claims.Username)
		c.Set("roles", claims.Roles)
		c.Next()
	}
}
