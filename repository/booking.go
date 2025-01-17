package repository

import (
	"context"
	"errors"

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
func (r *BookingRepository) GetUserID(id uint) (uint, error) {
	var booking entity.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("booking not found")
		}
		return 0, err
	}
	return booking.UserID, nil
}
func (r *BookingRepository) GetManyByUser(ctx context.Context, userID uint) ([]interface{}, error) {

	bookings, err := r.GetBookingsByUser(ctx, userID)
	if err != nil {

		return nil, err

	}
	result := make([]interface{}, len(bookings))
	for i, booking := range bookings {

		result[i] = booking

	}
	return result, nil

}

func (r *BookingRepository) GetBookingsByUser(ctx context.Context, userId uint) ([]*entity.Booking, error) {
	var bookings []*entity.Booking
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
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
	res := r.db.Model(&booking).Preload("User").Preload("Schedule").Preload("Payment").Preload("Cancellation").Preload("TicketUsage").Preload("Refund").Where("id = ?", bookingId).First(&booking)

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
