package entity

import "time"

type Cancellation struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	BookingID          uint      `json:"booking_id"`
	CancellationReason string    `json:"cancellation_reason"`
	CancellationDate   time.Time `json:"cancellation_date" gorm:"default:CURRENT_TIMESTAMP"`
	Booking            Booking   `json:"booking" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}

type CancellationRepository interface{}
