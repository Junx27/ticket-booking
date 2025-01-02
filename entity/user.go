package entity

import (
	"context"
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	FirsName    string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role" gorm:"default:customer"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserRepository interface {
	GetMany(ctx context.Context) ([]User, error)
}
