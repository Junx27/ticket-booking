package entity

import "time"

type TicketUsage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BookingID uint      `json:"booking_id" gorm:"not null"`
	IsUsed    bool      `json:"is_used" gorm:"default:false"`
	UsedAt    time.Time `json:"used_at"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Booking   Booking   `json:"booking" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}
