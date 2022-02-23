package persistence

import (
	"context"
	"coupon-service/domains/entities"
	"coupon-service/domains/repositories"

	"gorm.io/gorm"
)

type BoardUserRepositoryInstance struct {
	db *gorm.DB
}

func NewBoardUserRepository(db *gorm.DB) repositories.BoardUserRepository {
	return &BoardUserRepositoryInstance{
		db: db,
	}
}

func (repo *BoardUserRepositoryInstance) Create(ctx context.Context, newUser *entities.BoardUser) (*entities.BoardUser, error) {
	err := repo.db.Create(newUser).Error
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
