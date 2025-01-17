package repository

import (
	"context"
	"errors"

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

func (r *NotificationRepository) GetUserID(id uint) (uint, error) {
	var notification entity.Payment
	if err := r.db.First(&notification, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("notification not found")
		}
		return 0, err
	}
	return notification.UserID, nil
}
func (r *NotificationRepository) GetManyByUser(ctx context.Context, userID uint) ([]interface{}, error) {

	notifications, err := r.GetMany(ctx, userID)
	if err != nil {

		return nil, err

	}
	result := make([]interface{}, len(notifications))
	for i, notification := range notifications {

		result[i] = notification

	}
	return result, nil

}

func (r *NotificationRepository) GetMany(ctx context.Context, userId uint) ([]*entity.Notification, error) {
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
