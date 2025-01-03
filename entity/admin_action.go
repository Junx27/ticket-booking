package entity

import "time"

type AdminAction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	ActionType  string    `json:"action_type" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	ActionDate  time.Time `json:"action_date" gorm:"default:CURRENT_TIMESTAMP"`
	User        User      `json:"admin" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type AdminActionRepository interface{}
