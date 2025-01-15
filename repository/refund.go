package repository

import (
	"context"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type RefundRepository struct {
	db *gorm.DB
}

func NewRefundRepository(db *gorm.DB) entity.RefundRepository {
	return &RefundRepository{
		db: db,
	}
}

func (r *RefundRepository) GetMany(ctx context.Context) ([]*entity.Refund, error) {
	var refunds []*entity.Refund
	if err := r.db.WithContext(ctx).Find(&refunds).Error; err != nil {
		return nil, err
	}
	return refunds, nil
}

func (r *RefundRepository) GetOne(ctx context.Context, refundId uint) (*entity.RefundWithRelation, error) {
	refund := &entity.RefundWithRelation{}
	res := r.db.Model(&refund).Preload("Booking").Where("id = ?", refundId).First(&refund)

	if res.Error != nil {
		return nil, res.Error
	}

	return refund, nil
}

func (r *RefundRepository) CreateOne(ctx context.Context, refund *entity.Refund) (*entity.Refund, error) {
	if err := r.db.WithContext(ctx).Create(refund).Error; err != nil {
		return nil, err
	}

	return refund, nil
}

func (r *RefundRepository) UpdateOne(ctx context.Context, refundId uint, updateData map[string]interface{}) (*entity.Refund, error) {
	refund := &entity.Refund{}
	res := r.db.Model(&refund).Where("id = ?", refundId).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return refund, nil
}

func (r *RefundRepository) DeleteOne(ctx context.Context, refundId uint) error {
	refund := &entity.Refund{}
	res := r.db.Model(&refund).Where("id = ?", refundId).Delete(&refund)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
