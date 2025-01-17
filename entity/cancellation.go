package entity

import (
	"context"
	"time"
)

type BaseModelCancellation struct{}

func (BaseModelCancellation) TableName() string {
	return "cancellations"
}

type Cancellation struct {
	BaseModelCancellation
	ID                 uint               `json:"id" gorm:"primaryKey"`
	UserID             uint               `json:"user_id"`
	BookingID          uint               `json:"booking_id"`
	CancellationReason string             `json:"cancellation_reason"`
	CancellationDate   time.Time          `json:"cancellation_date" gorm:"default:CURRENT_TIMESTAMP"`
	User               UserDetailResponse `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Booking            Booking            `json:"-" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}
type CancellationWithRelation struct {
	BaseModelCancellation
	ID                 uint               `json:"id" gorm:"primaryKey"`
	UserID             uint               `json:"-"`
	BookingID          uint               `json:"booking_id"`
	CancellationReason string             `json:"cancellation_reason"`
	User               UserDetailResponse `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CancellationDate   time.Time          `json:"cancellation_date" gorm:"default:CURRENT_TIMESTAMP"`
	Booking            Booking            `json:"booking" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}

type CancellationRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*Cancellation, error)
	GetManyByBookingID(ctx context.Context, bookingId uint) ([]*Cancellation, error)
	GetOne(ctx context.Context, cancellationId uint) (*CancellationWithRelation, error)
	CreateOne(ctx context.Context, cancellation *Cancellation) (*Cancellation, error)
	DeleteOne(ctx context.Context, cancellationId uint) error
	DeleteMany(ctx context.Context) error
}
