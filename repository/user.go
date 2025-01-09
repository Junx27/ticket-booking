package repository

import (
	"context"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type UserRespository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &UserRespository{
		db: db,
	}
}

func (r *UserRespository) GetMany(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *UserRespository) GetOne(ctx context.Context, userId uint) (*entity.UserWithRelation, error) {
	user := &entity.UserWithRelation{}
	res := r.db.Model(&user).Where("id = ?", userId).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *UserRespository) UpdateOne(ctx context.Context, userId uint, updateData map[string]interface{}) (*entity.User, error) {
	user := &entity.User{}
	res := r.db.Model(&user).Where("id = ?", userId).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
func (r *UserRespository) DeleteOne(ctx context.Context, userId uint) error {
	user := &entity.User{}
	res := r.db.Model(&user).Where("id = ?", userId).Delete(&user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
