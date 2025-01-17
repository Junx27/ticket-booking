package entity

import (
	"context"
	"time"
)

type BaseModelRefund struct{}

func (BaseModelRefund) TableName() string {
	return "refunds"
}

type Refund struct {
	BaseModelRefund
	ID                uint      `json:"id" gorm:"primaryKey"`
	BookingID         uint      `json:"booking_id" gorm:"not null"`
	RefundAmount      float64   `json:"refund_amount" gorm:"not null"`
	RefundStatus      string    `json:"refund_status" gorm:"default:'pending'"`
	RefundDate        time.Time `json:"refund_date" gorm:"default:CURRENT_TIMESTAMP"`
	BankAccountNumber string    `json:"bank_account_number"`
	BankAccountName   string    `json:"bank_account_name"`
	CreatedAt         time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Booking           Booking   `json:"-" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}
type RefundWithRelation struct {
	BaseModelRefund
	ID                uint      `json:"id" gorm:"primaryKey"`
	BookingID         uint      `json:"-" gorm:"not null"`
	RefundAmount      float64   `json:"refund_amount" gorm:"not null"`
	RefundStatus      string    `json:"refund_status" gorm:"default:'pending'"`
	RefundDate        time.Time `json:"refund_date" gorm:"default:CURRENT_TIMESTAMP"`
	BankAccountNumber string    `json:"bank_account_number"`
	BankAccountName   string    `json:"bank_account_name"`
	CreatedAt         time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Booking           Booking   `json:"booking" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}

type RefundRepository interface {
	GetMany(ctx context.Context) ([]*Refund, error)
	GetOne(ctx context.Context, refundId uint) (*RefundWithRelation, error)
	CreateOne(ctx context.Context, refund *Refund) (*Refund, error)
	UpdateOne(ctx context.Context, refundId uint, updateData map[string]interface{}) (*Refund, error)
	DeleteOne(ctx context.Context, refundId uint) error
}
