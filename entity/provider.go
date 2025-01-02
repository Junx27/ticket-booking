package entity

import "time"

type Provider struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description"`
	ContactInfo string     `json:"contact_info"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Schedules   []Schedule `json:"schedules" gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE"`
}
