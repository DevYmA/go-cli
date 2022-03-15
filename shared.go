package main

import (
	"strings"
)

func validateInputs(firstName string, lastName string, email string) (bool, bool) {
	isValidFullName := len(firstName) > 4 && len(lastName) > 4
	isValidEmail := strings.Contains(email, "@")

	return isValidFullName, isValidEmail
}
