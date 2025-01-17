package entity

import (
	"context"
	"time"

	"github.com/lib/pq"
)

type BaseModelBooking struct{}

func (BaseModelBooking) TableName() string {
	return "bookings"
}

type Booking struct {
	BaseModelBooking
	ID            uint          `json:"id" gorm:"primaryKey"`
	UserID        uint          `json:"user_id"`
	ScheduleID    uint          `json:"schedule_id"`
	TicketCode    string        `json:"ticket_code" gorm:"not null"`
	TotalAmount   float64       `json:"total_amount" gorm:"not null"`
	BookingStatus string        `json:"booking_status" gorm:"default:pending"`
	SeatNumbers   pq.Int64Array `json:"seat_numbers" gorm:"type:integer[]"`
	CreatedAt     time.Time     `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time     `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type BookingWithRelation struct {
	BaseModelBooking
	ID            uint                 `json:"id" gorm:"primaryKey"`
	UserID        uint                 `json:"user_id"`
	ScheduleID    uint                 `json:"schedule_id"`
	TicketCode    string               `json:"ticket_code" gorm:"not null"`
	TotalAmount   float64              `json:"total_amount" gorm:"not null"`
	BookingStatus string               `json:"booking_status" gorm:"default:pending"`
	SeatNumbers   pq.Int64Array        `json:"seat_numbers" gorm:"type:integer[]"`
	CreatedAt     time.Time            `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time            `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	User          UserDetailResponse   `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Schedule      Schedule             `json:"schedule" gorm:"foreignKey:ScheduleID;constraint:OnDelete:CASCADE"`
	Payment       *PaymentWithRelation `json:"payment,omitempty" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	Cancellation  *Cancellation        `json:"cancellation,omitempty" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	TicketUsage   *TicketUsageStatus   `json:"ticket_usage,omitempty" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	Refund        *Refund              `json:"refund,omitempty" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}

type BookingRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*Booking, error)
	GetOne(ctx context.Context, bookingId uint) (*BookingWithRelation, error)
	CreateOne(ctx context.Context, booking *Booking) (*Booking, error)
	UpdateOne(ctx context.Context, bookingId uint, updateData map[string]interface{}) (*Booking, error)
	DeleteOne(ctx context.Context, bookingId uint) error
}
