package repository

import (
	"context"
	"errors"

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

func (r *PaymentRepository) GetUserID(id uint) (uint, error) {
	var payment entity.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("payment not found")
		}
		return 0, err
	}
	return payment.UserID, nil
}
func (r *PaymentRepository) GetManyByUser(ctx context.Context, userID uint) ([]interface{}, error) {

	payments, err := r.GetMany(ctx, userID)
	if err != nil {

		return nil, err

	}
	result := make([]interface{}, len(payments))
	for i, payment := range payments {

		result[i] = payment

	}
	return result, nil

}

func (r *PaymentRepository) GetMany(ctx context.Context, userId uint) ([]*entity.Payment, error) {
	var payments []*entity.Payment
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *PaymentRepository) GetManyProvider(ctx context.Context) ([]*entity.Payment, error) {
	var payments []*entity.Payment
	if err := r.db.WithContext(ctx).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *PaymentRepository) GetOne(ctx context.Context, paymentId uint) (*entity.PaymentWithRelation, error) {
	payment := &entity.PaymentWithRelation{}
	if err := r.db.WithContext(ctx).Preload("Booking").Where("id = ?", paymentId).First(&payment).Error; err != nil {
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
