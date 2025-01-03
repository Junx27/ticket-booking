package entity

import "time"

type Notification struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	UserID           uint      `json:"user_id"`
	NotificationType string    `json:"notification_type" gorm:"not null"`
	Message          string    `json:"message" gorm:"not null"`
	Status           string    `json:"status" gorm:"default:unread"`
	CreatedAt        time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	User             User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type NotificationRepository interface{}
