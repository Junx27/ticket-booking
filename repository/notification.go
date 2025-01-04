package repository

import (
	"context"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) entity.NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) GetMany(ctx context.Context) ([]*entity.Notification, error) {
	var notifications []*entity.Notification
	if err := r.db.WithContext(ctx).Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *NotificationRepository) GetManyByUser(ctx context.Context, userId uint) ([]*entity.Notification, error) {
	var notifications []*entity.Notification
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *NotificationRepository) GetOne(ctx context.Context, notificationId uint) (*entity.Notification, error) {
	notification := &entity.Notification{}
	if err := r.db.WithContext(ctx).Where("id = ?", notificationId).First(&notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *NotificationRepository) CreateOne(ctx context.Context, notification *entity.Notification) (*entity.Notification, error) {
	if err := r.db.WithContext(ctx).Create(notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *NotificationRepository) UpdateOne(ctx context.Context, notificationId uint, updateData map[string]interface{}) (*entity.Notification, error) {
	notification := &entity.Notification{}
	if err := r.db.WithContext(ctx).Where("id = ?", notificationId).First(&notification).Error; err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Model(&notification).Updates(updateData).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *NotificationRepository) DeleteOne(ctx context.Context, notificationId uint) error {
	notification := &entity.Notification{}
	if err := r.db.WithContext(ctx).Where("id = ?", notificationId).First(&notification).Error; err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).Delete(&notification).Error; err != nil {
		return err
	}
	return nil
}

func (r *NotificationRepository) DeleteAllByUser(ctx context.Context, userId uint) error {
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Delete(&entity.Notification{}).Error; err != nil {
		return err
	}
	return nil
}
