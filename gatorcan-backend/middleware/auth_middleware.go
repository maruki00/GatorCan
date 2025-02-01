package middleware

import (
	"fmt"
	"gatorcan-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			fmt.Println("No Authorization token received")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

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
		c.Set("role", claims.Role)

		c.Next()
	}
}
