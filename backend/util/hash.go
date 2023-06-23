package util

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"inventory-management/backend/internal/http/presenter/response"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New(response.InvalidPassword)
	}

	return nil
}
