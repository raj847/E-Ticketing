package repository

import (
	"context"
	"eticketing/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Register(ctx context.Context, user *entity.User) error {
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var res entity.User
	err := u.db.WithContext(ctx).Table("users").Where("username = ?", username).Find(&res).Error
	if err != nil {
		return entity.User{}, err
	}

	return res, nil
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var res entity.User
	err := u.db.WithContext(ctx).Table("users").Where("email = ?", email).Find(&res).Error
	if err != nil {
		return entity.User{}, err
	}

	return res, nil
}
