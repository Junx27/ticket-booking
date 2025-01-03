package entity

import "time"

type Schedule struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	ProviderID        uint      `json:"provider_id"`
	DepartureTime     time.Time `json:"departure_time" gorm:"not null"`
	ArrivalTime       time.Time `json:"arrival_time" gorm:"not null"`
	DepartureLocation string    `json:"departure_location" gorm:"not null"`
	ArrivalLocation   string    `json:"arrival_location" gorm:"not null"`
	AvailableSeats    int       `json:"available_seats" gorm:"not null"`
	TicketPrice       float64   `json:"ticket_price" gorm:"not null"`
	CreatedAt         time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Provider          Provider  `json:"provider" gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE"`
	Bookings          []Booking `json:"bookings" gorm:"foreignKey:ScheduleID;constraint:OnDelete:CASCADE"`
}

type ScheduleRepository interface{}
