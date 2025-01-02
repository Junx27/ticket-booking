package entity

import (
	"context"
	"time"
)

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
type UserWithRelation struct {
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
	AdminActions  []AdminAction  `json:"admin_actions" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (UserWithRelation) TableName() string {
	return "users"
}

type UserRepository interface {
	GetMany(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, userId uint) (*UserWithRelation, error)
}
