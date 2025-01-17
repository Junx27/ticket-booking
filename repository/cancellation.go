package repository

import (
	"context"
	"errors"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type CancellationRepository struct {
	db *gorm.DB
}

func NewCancellationRepository(db *gorm.DB) entity.CancellationRepository {
	return &CancellationRepository{
		db: db,
	}
}

func (r *CancellationRepository) GetUserID(id uint) (uint, error) {
	var cancellation entity.Cancellation
	if err := r.db.First(&cancellation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("cancellation not found")
		}
		return 0, err
	}
	return cancellation.UserID, nil
}
func (r *CancellationRepository) GetManyByUser(ctx context.Context, userID uint) ([]interface{}, error) {

	cancellations, err := r.GetMany(ctx, userID)
	if err != nil {

		return nil, err

	}
	result := make([]interface{}, len(cancellations))
	for i, cancellation := range cancellations {

		result[i] = cancellation

	}
	return result, nil

}

func (r *CancellationRepository) GetMany(ctx context.Context, userId uint) ([]*entity.Cancellation, error) {
	var cancellations []*entity.Cancellation
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&cancellations).Error; err != nil {
		return nil, err
	}
	return cancellations, nil
}

func (r *CancellationRepository) GetManyByBookingID(ctx context.Context, bookingId uint) ([]*entity.Cancellation, error) {
	var cancellations []*entity.Cancellation
	if err := r.db.WithContext(ctx).Where("booking_id = ?", bookingId).Find(&cancellations).Error; err != nil {
		return nil, err
	}
	return cancellations, nil
}

func (r *CancellationRepository) GetOne(ctx context.Context, cancellationId uint) (*entity.CancellationWithRelation, error) {
	cancellation := &entity.CancellationWithRelation{}
	if err := r.db.WithContext(ctx).Preload("Booking").Where("id = ?", cancellationId).First(&cancellation).Error; err != nil {
		return nil, err
	}
	return cancellation, nil
}

func (r *CancellationRepository) CreateOne(ctx context.Context, cancellation *entity.Cancellation) (*entity.Cancellation, error) {
	if err := r.db.WithContext(ctx).Create(cancellation).Error; err != nil {
		return nil, err
	}
	return cancellation, nil
}

func (r *CancellationRepository) DeleteOne(ctx context.Context, cancellationId uint) error {
	cancellation := &entity.Cancellation{}
	if err := r.db.WithContext(ctx).Where("id = ?", cancellationId).First(&cancellation).Error; err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).Delete(&cancellation).Error; err != nil {
		return err
	}
	return nil
}

func (r *CancellationRepository) DeleteMany(ctx context.Context) error {
	if err := r.db.WithContext(ctx).Exec("TRUNCATE TABLE cancellations RESTART IDENTITY CASCADE").Error; err != nil {
		return err
	}
	return nil
}
