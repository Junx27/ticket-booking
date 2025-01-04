package entity

import (
	"context"
	"time"
)

type Payment struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	BookingID     uint      `json:"booking_id"`
	PaymentAmount float64   `json:"payment_amount" gorm:"not null"`
	PaymentStatus string    `json:"payment_status" gorm:"default:pending"`
	PaymentMethod string    `json:"payment_method" gorm:"not null"`
	PaymentDate   time.Time `json:"payment_date" gorm:"default:CURRENT_TIMESTAMP"`
	Booking       Booking   `json:"booking" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}

type PaymentRepository interface {
	GetMany(ctx context.Context) ([]*Payment, error)
	GetOne(ctx context.Context, paymentId uint) (*Payment, error)
	CreateOne(ctx context.Context, payment *Payment) (*Payment, error)
	UpdateOne(ctx context.Context, paymentId uint, updateData map[string]interface{}) (*Payment, error)
	DeleteOne(ctx context.Context, paymentId uint) error
}
