package entity

import (
	"context"
	"time"
)

type BaseModelProvider struct{}

func (BaseModelProvider) TableName() string {
	return "providers"
}

type Provider struct {
	BaseModelProvider
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	ContactInfo string    `json:"contact_info"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
type ProviderWithRelation struct {
	BaseModelProvider
	ID          uint       `json:"id" gorm:"primaryKey"`
	UserID      uint       `json:"user_id"`
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description"`
	ContactInfo string     `json:"contact_info"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Schedules   []Schedule `json:"schedules" gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE"`
}

type ProviderRepository interface {
	GetMany(ctx context.Context) ([]*Provider, error)
	GetManyByUser(ctx context.Context, userId uint) ([]*Provider, error)
	GetOne(ctx context.Context, providerId uint) (*ProviderWithRelation, error)
	CreateOne(ctx context.Context, provider *Provider) (*Provider, error)
	UpdateOne(ctx context.Context, providerId uint, updateData map[string]interface{}) (*Provider, error)
	DeleteOne(ctx context.Context, providerId uint) error
}
