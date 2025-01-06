package repository

import (
	"context"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) entity.BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) GetMany(ctx context.Context) ([]*entity.Booking, error) {
	var bookings []*entity.Booking
	if err := r.db.WithContext(ctx).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *BookingRepository) GetOne(ctx context.Context, bookingId uint) (*entity.BookingWithRelation, error) {
	booking := &entity.BookingWithRelation{}
	res := r.db.Model(&booking).Where("id = ?", bookingId).First(&booking)

	if res.Error != nil {
		return nil, res.Error
	}

	return booking, nil
}

func (r *BookingRepository) CreateOne(ctx context.Context, booking *entity.Booking) (*entity.Booking, error) {
	if err := r.db.WithContext(ctx).Create(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *BookingRepository) UpdateOne(ctx context.Context, bookingId uint, updateData map[string]interface{}) (*entity.Booking, error) {
	booking := &entity.Booking{}
	res := r.db.Model(&booking).Where("id = ?", bookingId).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return booking, nil
}

func (r *BookingRepository) DeleteOne(ctx context.Context, bookingId uint) error {
	booking := &entity.Booking{}
	res := r.db.Model(&booking).Where("id = ?", bookingId).Delete(&booking)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
