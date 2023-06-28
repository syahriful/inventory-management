package model

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	response "inventory-management/backend/internal/http/response"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := u.HashPassword()
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return nil
}

func (u *User) ToResponse() *response.UserResponse {
	return &response.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Username:  u.Username,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}

func (u *User) HashPassword(password ...string) (string, error) {
	if len(password) > 0 {
		u.Password = password[0]
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.New(response.ErrorInvalidPassword)
	}

	return nil
}

func (u *User) GenerateTokenJWT() (string, error) {
	myClaims := jwt.Claims(jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
