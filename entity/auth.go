package entity

import (
	"context"
	"net/mail"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type AuthCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, registerData *User) (*User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User, error)
}

type AuthService interface {
	Login(ctx context.Context, loginData *AuthCredentials) (string, *User, error)
	Register(ctx context.Context, registerData *User) (string, *User, error)
	Logout(ctx context.Context, token string) error
}

func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	reUppercase := regexp.MustCompile(`[A-Z]`)
	reDigit := regexp.MustCompile(`\d`)
	reSymbol := regexp.MustCompile(`[\W_]`)
	return reUppercase.MatchString(password) && reDigit.MatchString(password) && reSymbol.MatchString(password)
}
