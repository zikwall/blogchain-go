package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func BlogchainPasswordCorrectness(hash string, password string) bool {
	errf := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}

	return true
}

func GenerateBlogchainPasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
