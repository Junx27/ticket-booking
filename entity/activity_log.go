package entity

import (
	"context"
	"time"
)

type BaseModelActivityLog struct{}

func (BaseModelActivityLog) TableName() string {
	return "activity_logs"
}

type ActivityLog struct {
	BaseModelActivityLog
	ID          uint               `json:"id" gorm:"primaryKey"`
	UserID      uint               `json:"user_id"`
	ActionType  string             `json:"action_type" gorm:"default:INFO"`
	Description string             `json:"description" gorm:"not null"`
	Timestamp   time.Time          `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP"`
	User        UserDetailResponse `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
type ActivityLogWithRelation struct {
	BaseModelActivityLog
	ID          uint               `json:"id" gorm:"primaryKey"`
	UserID      uint               `json:"user_id"`
	ActionType  string             `json:"action_type" gorm:"default:INFO"`
	Description string             `json:"description" gorm:"not null"`
	Timestamp   time.Time          `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP"`
	User        UserDetailResponse `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type ActivityLogRepository interface {
	GetMany(ctx context.Context) ([]*ActivityLogWithRelation, error)
	GetManyByUser(ctx context.Context, userId uint) ([]*ActivityLog, error)
	GetOne(ctx context.Context, activityLogId uint) (*ActivityLogWithRelation, error)
	CreateOne(ctx context.Context, activityLog *ActivityLog) (*ActivityLog, error)
	DeleteOne(ctx context.Context, activityLogId uint) error
	DeleteMany(ctx context.Context) error
}
