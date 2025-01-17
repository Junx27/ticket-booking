package entity

import (
	"context"
	"time"

	"github.com/lib/pq"
)

type BaseModelSchedule struct{}

func (BaseModelSchedule) TableName() string {
	return "schedules"
}

type Schedule struct {
	BaseModelSchedule
	ID                uint               `json:"id" gorm:"primaryKey"`
	UserID            uint               `json:"user_id"`
	ProviderID        uint               `json:"provider_id"`
	DepartureTime     time.Time          `json:"departure_time" gorm:"not null"`
	ArrivalTime       time.Time          `json:"arrival_time" gorm:"not null"`
	DepartureLocation string             `json:"departure_location" gorm:"not null"`
	ArrivalLocation   string             `json:"arrival_location" gorm:"not null"`
	AvailableSeats    pq.Int64Array      `json:"available_seats" gorm:"type:integer[]"`
	TicketPrice       float64            `json:"ticket_price" gorm:"not null"`
	User              UserDetailResponse `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Provider          Provider           `json:"-" gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE"`
	Bookings          []Booking          `json:"-" gorm:"foreignKey:ScheduleID;constraint:OnDelete:CASCADE"`
}
type ScheduleWithRelation struct {
	BaseModelSchedule
	ID                uint          `json:"id" gorm:"primaryKey"`
	UserID            uint          `json:"user_id"`
	ProviderID        uint          `json:"provider_id"`
	DepartureTime     time.Time     `json:"departure_time" gorm:"not null"`
	ArrivalTime       time.Time     `json:"arrival_time" gorm:"not null"`
	DepartureLocation string        `json:"departure_location" gorm:"not null"`
	ArrivalLocation   string        `json:"arrival_location" gorm:"not null"`
	AvailableSeats    pq.Int64Array `json:"available_seats" gorm:"type:integer[]"`
	TicketPrice       float64       `json:"ticket_price" gorm:"not null"`
	CreatedAt         time.Time     `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time     `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Provider          Provider      `json:"provider" gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE"`
	Bookings          []Booking     `json:"bookings" gorm:"foreignKey:ScheduleID;constraint:OnDelete:CASCADE"`
}

type ScheduleRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*Schedule, error)
	GetOne(ctx context.Context, scheduleId uint) (*Schedule, error)
	CreateOne(ctx context.Context, schedule *Schedule) (*Schedule, error)
	UpdateOne(ctx context.Context, scheduleId uint, updateData map[string]interface{}) (*Schedule, error)
	UpdateSeatsStatus(ctx context.Context, scheduleId uint, seatsData map[int]interface{}) (*Schedule, error)
	DeleteOne(ctx context.Context, scheduleId uint) error
}

func (s *Schedule) IsSeatAvailable(seatNumber int64) bool {
	for _, seat := range s.AvailableSeats {
		if seat == seatNumber {
			return true
		}
	}
	return false
}
