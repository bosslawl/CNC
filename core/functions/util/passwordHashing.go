package Util

import (
	"golang.org/x/crypto/bcrypt"
)

// takes current password then SHA256 SUM Hashes it and returns the outcome
func PasswordHash(password string) string {
	pwd := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
