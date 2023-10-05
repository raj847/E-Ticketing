package repository

import (
	"context"
	"eticketing/entity"

	"gorm.io/gorm"
)

type TerminalRepository struct {
	db *gorm.DB
}

func NewTerminalRepository(db *gorm.DB) *TerminalRepository {
	return &TerminalRepository{db}
}

func (u *TerminalRepository) CreateTerminal(ctx context.Context, terminal *entity.Terminal) error {
	err := u.db.WithContext(ctx).Create(&terminal).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *TerminalRepository) GetTerminal(ctx context.Context) ([]entity.Terminal, error) {
	var res []entity.Terminal
	err := u.db.WithContext(ctx).Table("terminals").Find(&res).Error
	if err != nil {
		return []entity.Terminal{}, err
	}

	return res, nil
}
