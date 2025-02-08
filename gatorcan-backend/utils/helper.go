package utils

import "github.com/gin-gonic/gin"

func HasRole(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func IsAdmin(c *gin.Context) bool {
	roles, exists := c.Get("roles")
	if !exists {
		return false
	}
	rolesSlice, ok := roles.([]string)
	if !ok {
		return false
	}
	return HasRole(rolesSlice, "admin")
}
