package entity

import (
	"context"
	"time"
)

type Notification struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	UserID           uint      `json:"user_id"`
	NotificationType string    `json:"notification_type" gorm:"default:INFO"`
	Message          string    `json:"message" gorm:"not null"`
	Status           string    `json:"status" gorm:"default:unread"`
	CreatedAt        time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type NotificationRepository interface {
	GetMany(ctx context.Context) ([]*Notification, error)
	GetManyByUser(ctx context.Context, userId uint) ([]*Notification, error)
	GetOne(ctx context.Context, notificationId uint) (*Notification, error)
	CreateOne(ctx context.Context, notification *Notification) (*Notification, error)
	UpdateOne(ctx context.Context, notificationId uint, updateData map[string]interface{}) (*Notification, error)
	DeleteOne(ctx context.Context, notificationId uint) error
	DeleteAllByUser(ctx context.Context, userId uint) error
}
