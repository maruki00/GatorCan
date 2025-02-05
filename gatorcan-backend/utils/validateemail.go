package utils

import "regexp"

func IsValidEmail(email string) bool {
	// Regular expression for validating email format
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
