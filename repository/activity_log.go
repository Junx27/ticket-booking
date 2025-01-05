package repository

import (
	"context"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type ActivityLogRepository struct {
	db *gorm.DB
}

func NewActivityLogRepository(db *gorm.DB) entity.ActivityLogRepository {
	return &ActivityLogRepository{
		db: db,
	}
}

func (r *ActivityLogRepository) GetMany(ctx context.Context) ([]*entity.ActivityLog, error) {
	var activityLogs []*entity.ActivityLog
	if err := r.db.WithContext(ctx).Find(&activityLogs).Error; err != nil {
		return nil, err
	}
	return activityLogs, nil
}

func (r *ActivityLogRepository) GetManyByUser(ctx context.Context, userId uint) ([]*entity.ActivityLog, error) {
	var activityLogs []*entity.ActivityLog
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&activityLogs).Error; err != nil {
		return nil, err
	}
	return activityLogs, nil
}

func (r *ActivityLogRepository) GetOne(ctx context.Context, activityLogId uint) (*entity.ActivityLog, error) {
	activityLogs := &entity.ActivityLog{}
	if err := r.db.WithContext(ctx).Where("id = ?", activityLogId).First(&activityLogs).Error; err != nil {
		return nil, err
	}
	return activityLogs, nil
}

func (r *ActivityLogRepository) CreateOne(ctx context.Context, activityLog *entity.ActivityLog) (*entity.ActivityLog, error) {
	if err := r.db.WithContext(ctx).Create(activityLog).Error; err != nil {
		return nil, err
	}
	return activityLog, nil
}

func (r *ActivityLogRepository) DeleteOne(ctx context.Context, activityLogId uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.ActivityLog{}, activityLogId).Error; err != nil {
		return err
	}
	return nil
}

func (r *ActivityLogRepository) DeleteMany(ctx context.Context) error {
	if err := r.db.WithContext(ctx).Exec("TRUNCATE TABLE activity_logs RESTART IDENTITY CASCADE").Error; err != nil {
		return err
	}
	return nil
}
