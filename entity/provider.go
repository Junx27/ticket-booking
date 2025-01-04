package entity

import (
	"context"
	"time"
)

type Provider struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description"`
	ContactInfo string     `json:"contact_info"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Schedules   []Schedule `json:"schedules" gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE"`
}

type ProviderRepository interface {
	GetMany(ctx context.Context) ([]*Provider, error)
	GetOne(ctx context.Context, providerId uint) (*Provider, error)
	CreateOne(ctx context.Context, provider *Provider) (*Provider, error)
	UpdateOne(ctx context.Context, providerId uint, updateData map[string]interface{}) (*Provider, error)
	DeleteOne(ctx context.Context, providerId uint) error
}