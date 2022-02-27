package repositories

import (
	"context"
	"coupon-service/domains/entities"

	"gorm.io/gorm"
)

type BoardUserRepository interface {
	Find(ctx context.Context) ([]entities.BoardUser, error)
	Create(ctx context.Context, newUser *entities.BoardUser) (*entities.BoardUser, error)
}

type BoardUserRepositoryInstance struct {
	db *gorm.DB
}

func NewBoardUserRepository(db *gorm.DB) BoardUserRepository {
	return &BoardUserRepositoryInstance{
		db: db,
	}
}

func (repo *BoardUserRepositoryInstance) Find(ctx context.Context) ([]entities.BoardUser, error) {
	foundUsers := make([]entities.BoardUser, 0)
	err := repo.db.WithContext(ctx).Find(&foundUsers).Error
	if err != nil {
		return nil, err
	}

	return foundUsers, nil
}

func (repo *BoardUserRepositoryInstance) Create(ctx context.Context, newUser *entities.BoardUser) (*entities.BoardUser, error) {
	err := repo.db.WithContext(ctx).Create(newUser).Error
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
