package repository

import (
	"context"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) entity.PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) GetMany(ctx context.Context) ([]*entity.Payment, error) {
	var payments []*entity.Payment
	if err := r.db.WithContext(ctx).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *PaymentRepository) GetOne(ctx context.Context, paymentId uint) (*entity.Payment, error) {
	payment := &entity.Payment{}
	if err := r.db.WithContext(ctx).Where("id = ?", paymentId).First(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepository) CreateOne(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {
	if err := r.db.WithContext(ctx).Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepository) UpdateOne(ctx context.Context, paymentId uint, updateData map[string]interface{}) (*entity.Payment, error) {
	payment := &entity.Payment{}
	if err := r.db.WithContext(ctx).Where("id = ?", paymentId).First(&payment).Error; err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Model(&payment).Updates(updateData).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepository) DeleteOne(ctx context.Context, paymentId uint) error {
	payment := &entity.Payment{}
	if err := r.db.WithContext(ctx).Where("id = ?", paymentId).First(&payment).Error; err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).Delete(&payment).Error; err != nil {
		return err
	}
	return nil
}
