package repository

import (
	"context"

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

func (r *CancellationRepository) GetMany(ctx context.Context) ([]*entity.Cancellation, error) {
	var cancellations []*entity.Cancellation
	if err := r.db.WithContext(ctx).Find(&cancellations).Error; err != nil {
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
