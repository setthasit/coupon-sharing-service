package repositories

import (
	"context"
	"coupon-service/domains/entities"
	"coupon-service/infrastructure/errors"

	"gorm.io/gorm"
)

type BoardMemberRepository interface {
	ValidateUserInBoard(ctx context.Context, boardID uint, boardUserID uint) error
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

func (repo *BoardMemberRepositoryInstance) ValidateUserInBoard(ctx context.Context, boardID uint, boardUserID uint) error {
	member := new(entities.BoardMember)
	err := repo.db.WithContext(ctx).Where(entities.BoardMember{BoardID: boardID, BoardUserID: boardUserID}).First(member).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrBoardMemberNotMember()
		}
		return err
	}

	if member.BoardID == 0 || member.BoardUserID == 0 {
		return errors.ErrBoardMemberNotMember()
	}

	return nil
}

func (repo *BoardMemberRepositoryInstance) CreateBatch(ctx context.Context, newMembers []entities.BoardMember) ([]entities.BoardMember, error) {
	err := repo.db.WithContext(ctx).CreateInBatches(newMembers, 100).Error
	if err != nil {
		return nil, errors.ErrBoardUserCreateFailed()
	}

	return newMembers, nil
}
