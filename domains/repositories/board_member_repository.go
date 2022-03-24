package repositories

import (
	"context"
	"coupon-service/domains/entities"
	"coupon-service/infrastructure/errors"

	"gorm.io/gorm"
)

type BoardMemberRepository interface {
	CreateBatch(ctx context.Context, newMembers []entities.BoardMember) ([]entities.BoardMember, error)
}

type BoardMemberRepositoryInstance struct {
	db *gorm.DB
}

func NewBoardMemberRepository(db *gorm.DB) BoardMemberRepository {
	return &BoardMemberRepositoryInstance{
		db: db,
	}
}

func (repo *BoardMemberRepositoryInstance) CreateBatch(ctx context.Context, newMembers []entities.BoardMember) ([]entities.BoardMember, error) {
	err := repo.db.WithContext(ctx).CreateInBatches(newMembers, 100).Error
	if err != nil {
		return nil, errors.ErrBoardUserCreateFailed()
	}

	return newMembers, nil
}
