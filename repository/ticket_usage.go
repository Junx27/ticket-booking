package repository

import (
	"context"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type TicketUsageRepository struct {
	db *gorm.DB
}

func NewTicketUsageRepository(db *gorm.DB) entity.TicketUsageRepository {
	return &TicketUsageRepository{
		db: db,
	}
}

func (r *TicketUsageRepository) GetMany(ctx context.Context) ([]*entity.TicketUsage, error) {
	var ticketUsages []*entity.TicketUsage
	if err := r.db.WithContext(ctx).Find(&ticketUsages).Error; err != nil {
		return nil, err
	}
	return ticketUsages, nil
}

func (r *TicketUsageRepository) GetOne(ctx context.Context, ticketUsageId uint) (*entity.TicketUsage, error) {
	ticketUsage := &entity.TicketUsage{}
	res := r.db.Model(&ticketUsage).Where("id = ?", ticketUsageId).First(&ticketUsage)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticketUsage, nil
}

func (r *TicketUsageRepository) CreateOne(ctx context.Context, ticketUsage *entity.TicketUsage) (*entity.TicketUsage, error) {
	if err := r.db.WithContext(ctx).Create(ticketUsage).Error; err != nil {
		return nil, err
	}

	return ticketUsage, nil
}

func (r *TicketUsageRepository) UpdateOne(ctx context.Context, ticketUsageId uint, updateData map[string]interface{}) (*entity.TicketUsage, error) {
	ticketUsage := &entity.TicketUsage{}
	res := r.db.Model(&ticketUsage).Where("id = ?", ticketUsageId).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticketUsage, nil
}

func (r *TicketUsageRepository) DeleteOne(ctx context.Context, ticketUsageId uint) error {
	ticketUsage := &entity.TicketUsage{}
	res := r.db.Model(&ticketUsage).Where("id = ?", ticketUsageId).Delete(&ticketUsage)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
