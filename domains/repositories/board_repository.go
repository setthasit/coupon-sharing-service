package repositories

import (
	"context"
	"coupon-service/domains/entities"
	"coupon-service/infrastructure/errors"

	"gorm.io/gorm"
)

type BoardRepository interface {
	Find(ctx context.Context, boardUserID uint) ([]entities.Board, error)
	Create(ctx context.Context, newUser *entities.Board) (*entities.Board, error)
}

type BoardRepositoryInstance struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) BoardRepository {
	return &BoardRepositoryInstance{
		db: db,
	}
}

func (repo *BoardRepositoryInstance) Find(ctx context.Context, boardUserID uint) ([]entities.Board, error) {
	foundBoards := make([]entities.Board, 0)
	err := repo.db.
		WithContext(ctx).
		Preload("BoardMember", repo.db.Where(&entities.BoardMember{BoardUserID: boardUserID})).
		Find(&foundBoards).
		Error
	if err != nil {
		return nil, errors.NewAPIError(500, err)
	}

	return foundBoards, nil
}

func (repo *BoardRepositoryInstance) GetInfo(ctx context.Context, boardID uint, boardUserID uint) (*entities.Board, error) {
	foundBoard := new(entities.Board)
	err := repo.db.
		WithContext(ctx).
		Preload("BoardMember", repo.db.Where(&entities.BoardMember{BoardUserID: boardUserID})).
		Preload("Coupons", repo.db.Where(&entities.Coupon{BoardID: boardID})).
		First(&foundBoard).
		Error
	if err != nil {
		return nil, errors.NewAPIError(500, err)
	}

	return foundBoard, nil
}

func (repo *BoardRepositoryInstance) Create(ctx context.Context, newBoard *entities.Board) (*entities.Board, error) {
	err := repo.db.WithContext(ctx).Create(newBoard).Error
	if err != nil {
		return nil, errors.ErrBoardUserCreateFailed()
	}

	return newBoard, nil
}
