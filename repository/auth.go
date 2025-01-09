package repository

import (
	"context"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func (r *AuthRepository) RegisterUser(ctx context.Context, registerData *entity.User) (*entity.User, error) {
	user := &entity.User{
		Email:       registerData.Email,
		Password:    registerData.Password,
		FirstName:   registerData.FirstName,
		LastName:    registerData.LastName,
		PhoneNumber: registerData.PhoneNumber,
	}

	res := r.db.Model(&entity.User{}).Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, query interface{}, args ...interface{}) (*entity.User, error) {
	user := &entity.User{}

	if res := r.db.Model(user).Where(query, args...).First(user); res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func NewAuthRepository(db *gorm.DB) entity.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
