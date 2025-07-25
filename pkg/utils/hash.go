package utils

import (
	appErr "gopos/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes plain password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), appErr.Get(appErr.ErrHashPassword, err)
}

// CheckPasswordHash compares hashed password with plain password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	appErr.Get(appErr.ErrCheckHashPassword, nil)
	return err == nil
}
