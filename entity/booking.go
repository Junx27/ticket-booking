package entity

import "time"

type Booking struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserID          uint           `json:"user_id"`
	ScheduleID      uint           `json:"schedule_id"`
	NumberOfTickets int            `json:"number_of_tickets" gorm:"not null"`
	TotalAmount     float64        `json:"total_amount" gorm:"not null"`
	BookingStatus   string         `json:"booking_status" gorm:"default:pending"`
	CreatedAt       time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	User            User           `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Schedule        Schedule       `json:"schedule" gorm:"foreignKey:ScheduleID;constraint:OnDelete:CASCADE"`
	Payments        []Payment      `json:"payments" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	Cancellations   []Cancellation `json:"cancellations" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}
