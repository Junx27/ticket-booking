package entity

import (
	"context"
	"time"
)

type BaseModelTicketUsage struct{}

func (BaseModelTicketUsage) TableName() string {
	return "ticket_usages"
}

type TicketUsage struct {
	BaseModelTicketUsage
	ID        uint               `json:"id" gorm:"primaryKey"`
	UserID    uint               `json:"user_id"`
	BookingID uint               `json:"booking_id" gorm:"not null"`
	IsUsed    bool               `json:"is_used" gorm:"default:false"`
	UsedAt    time.Time          `json:"used_at"`
	CreatedAt time.Time          `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	User      UserDetailResponse `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type TicketUsageStatus struct {
	BaseModelTicketUsage
	ID        uint               `json:"id" gorm:"primaryKey"`
	UserID    uint               `json:"-"`
	BookingID uint               `json:"-" gorm:"not null"`
	IsUsed    bool               `json:"is_used" gorm:"default:false"`
	User      UserDetailResponse `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type TicketUsageRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*TicketUsage, error)
	GetOne(ctx context.Context, ticketUsageId uint) (*TicketUsage, error)
	CreateOne(ctx context.Context, ticketUsage *TicketUsage) (*TicketUsage, error)
	UpdateOne(ctx context.Context, ticketUsageId uint, updateData map[string]interface{}) (*TicketUsage, error)
	DeleteOne(ctx context.Context, ticketUsageId uint) error
}
