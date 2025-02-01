package utils

import (
	"fmt"
)

func GenerateadminToken() {
	username := "adminuser"
	role := "admin"

	token, err := GenerateToken(username, role)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated JWT Token:", token)
}
