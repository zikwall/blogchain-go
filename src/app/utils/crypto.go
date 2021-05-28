package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func BlogchainPasswordCorrectness(hash string, password string) bool {
	errf := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if errf != nil && errors.Is(errf, bcrypt.ErrMismatchedHashAndPassword) {
		return false
	}

	return true
}

func GenerateBlogchainPasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
