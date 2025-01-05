package entity

import (
	"context"
	"time"
)

type ActivityLog struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	ActionType  string    `json:"action_type" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Timestamp   time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP"`
	User        User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type ActivityLogRepository interface {
	GetMany(ctx context.Context) ([]*ActivityLog, error)
	GetManyByUser(ctx context.Context, userId uint) ([]*ActivityLog, error)
	GetOne(ctx context.Context, activityLogId uint) (*ActivityLog, error)
	CreateOne(ctx context.Context, activityLog *ActivityLog) (*ActivityLog, error)
	DeleteOne(ctx context.Context, activityLogId uint) error
	DeleteMany(ctx context.Context) error
}
