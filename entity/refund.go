package entity

import "time"

type Refund struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	BookingID         uint      `json:"booking_id" gorm:"not null"`
	RefundAmount      float64   `json:"refund_amount" gorm:"not null"`
	RefundStatus      string    `json:"refund_status" gorm:"default:'pending'"`
	RefundDate        time.Time `json:"refund_date" gorm:"default:CURRENT_TIMESTAMP"`
	BankAccountNumber string    `json:"bank_account_number"`
	BankAccountName   string    `json:"bank_account_name"`
	CreatedAt         time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Booking           Booking   `json:"booking" gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}
