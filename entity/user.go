package entity

import (
	"context"
	"time"
)

type BaseModelUser struct{}

func (BaseModelUser) TableName() string {
	return "users"
}

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	FirstName   string    `json:"first_name" gorm:"not null"`
	LastName    string    `json:"last_name" gorm:"not null"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role" gorm:"default:customer"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
type UserDetailResponse struct {
	BaseModelUser
	ID          uint   `json:"-" gorm:"primaryKey"`
	Email       string `json:"email" gorm:"unique;not null"`
	FirstName   string `json:"first_name" gorm:"not null"`
	PhoneNumber string `json:"phone_number"`
}

type UserWithRelation struct {
	BaseModelUser
	ID            uint           `json:"id" gorm:"primaryKey"`
	Email         string         `json:"email" gorm:"unique;not null"`
	Password      string         `json:"password" gorm:"not null"`
	FirstName     string         `json:"first_name" gorm:"not null"`
	LastName      string         `json:"last_name" gorm:"not null"`
	PhoneNumber   string         `json:"phone_number"`
	Role          string         `json:"role" gorm:"default:customer"`
	CreatedAt     time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Bookings      []Booking      `json:"bookings" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	ActivityLogs  []ActivityLog  `json:"activity_logs" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type UserRepository interface {
	GetMany(ctx context.Context) ([]*User, error)
	GetOne(ctx context.Context, userId uint) (*UserWithRelation, error)
	CreateOne(ctx context.Context, user *User) (*User, error)
	UpdateOne(ctx context.Context, userId uint, updateData map[string]interface{}) (*User, error)
	DeleteOne(ctx context.Context, userId uint) error
}
