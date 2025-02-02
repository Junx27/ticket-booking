package repository

import (
	"context"
	"errors"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type ProviderRepository struct {
	db *gorm.DB
}

func NewProviderRepository(db *gorm.DB) entity.ProviderRepository {
	return &ProviderRepository{
		db: db,
	}
}

func (r *ProviderRepository) GetUserID(id uint) (uint, error) {
	var provider entity.Provider
	if err := r.db.First(&provider, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("provider not found")
		}
		return 0, err
	}
	return provider.UserID, nil
}
func (r *ProviderRepository) GetManyByUser(ctx context.Context, userID uint) ([]interface{}, error) {

	providers, err := r.GetMany(ctx, userID)
	if err != nil {

		return nil, err

	}
	result := make([]interface{}, len(providers))
	for i, provider := range providers {

		result[i] = provider

	}
	return result, nil

}

func (r *ProviderRepository) GetMany(ctx context.Context, userId uint) ([]*entity.Provider, error) {
	var providers []*entity.Provider
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&providers).Error; err != nil {
		return nil, err
	}
	return providers, nil
}
func (r *ProviderRepository) GetManyCustomer(ctx context.Context) ([]*entity.Provider, error) {
	var providers []*entity.Provider
	if err := r.db.WithContext(ctx).Find(&providers).Error; err != nil {
		return nil, err
	}
	return providers, nil
}

func (r *ProviderRepository) GetOne(ctx context.Context, providerId uint) (*entity.ProviderWithRelation, error) {
	provider := &entity.ProviderWithRelation{}
	if err := r.db.WithContext(ctx).Preload("Schedules").Where("id = ?", providerId).First(&provider).Error; err != nil {
		return nil, err
	}

	return provider, nil
}

func (r *ProviderRepository) CreateOne(ctx context.Context, provider *entity.Provider) (*entity.Provider, error) {
	if err := r.db.WithContext(ctx).Create(provider).Error; err != nil {
		return nil, err
	}
	return provider, nil
}

func (r *ProviderRepository) UpdateOne(ctx context.Context, providerId uint, updateData map[string]interface{}) (*entity.Provider, error) {
	provider := &entity.Provider{}
	res := r.db.Model(&provider).Where("id = ?", providerId).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return provider, nil
}

func (r *ProviderRepository) DeleteOne(ctx context.Context, providerId uint) error {

	provider := &entity.Provider{}
	res := r.db.Model(&provider).Where("id = ?", providerId).Delete(&provider)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
