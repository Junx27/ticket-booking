package entity

import "time"

type ActivityLog struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	ActionType  string    `json:"action_type" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Timestamp   time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP"`
	User        User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type ActivityLogRepository interface{}
