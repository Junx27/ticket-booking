package entity

import (
	"context"
	"time"
)

type BaseModelPayment struct{}

func (BaseModelPayment) TableName() string {
	return "payments"
}

type Payment struct {
	BaseModelPayment
	ID            uint               `json:"id" gorm:"primaryKey"`
	UserID        uint               `json:"user_id"`
	BookingID     uint               `json:"booking_id"`
	PaymentAmount float64            `json:"payment_amount" gorm:"not null"`
	PaymentStatus string             `json:"payment_status" gorm:"default:success"`
	PaymentMethod string             `json:"payment_method" gorm:"not null"`
	PaymentDate   time.Time          `json:"payment_date" gorm:"default:CURRENT_TIMESTAMP"`
	User          UserDetailResponse `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Booking       Booking            `json:"-" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}
type PaymentWithRelation struct {
	BaseModelPayment
	ID            uint               `json:"id" gorm:"primaryKey"`
	UserID        uint               `json:"-"`
	BookingID     uint               `json:"-"`
	PaymentAmount float64            `json:"payment_amount" gorm:"not null"`
	PaymentStatus string             `json:"payment_status" gorm:"default:success"`
	PaymentMethod string             `json:"payment_method" gorm:"not null"`
	PaymentDate   time.Time          `json:"payment_date" gorm:"default:CURRENT_TIMESTAMP"`
	Booking       Booking            `json:"-" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	User          UserDetailResponse `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type PaymentRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*Payment, error)
	GetOne(ctx context.Context, paymentId uint) (*PaymentWithRelation, error)
	CreateOne(ctx context.Context, payment *Payment) (*Payment, error)
	UpdateOne(ctx context.Context, paymentId uint, updateData map[string]interface{}) (*Payment, error)
	DeleteOne(ctx context.Context, paymentId uint) error
}
