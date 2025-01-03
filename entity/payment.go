package entity

import "time"

type Payment struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	BookingID     uint      `json:"booking_id"`
	PaymentAmount float64   `json:"payment_amount" gorm:"not null"`
	PaymentStatus string    `json:"payment_status" gorm:"default:pending"`
	PaymentMethod string    `json:"payment_method" gorm:"not null"`
	PaymentDate   time.Time `json:"payment_date" gorm:"default:CURRENT_TIMESTAMP"`
	Booking       Booking   `json:"booking" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}

type PaymentRepository interface{}
