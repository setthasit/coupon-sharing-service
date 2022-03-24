package repositories

import (
	"context"
	"coupon-service/domains/entities"
	"coupon-service/infrastructure/errors"

	"gorm.io/gorm"
)

type BoardUserRepository interface {
	Find(ctx context.Context) ([]entities.BoardUser, error)
	FindByGoogleUserID(ctx context.Context, id string) (*entities.BoardUser, error)
	FindByEmail(ctx context.Context, email string) (*entities.BoardUser, error)
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

func (repo *BoardUserRepositoryInstance) FindByGoogleUserID(ctx context.Context, id string) (*entities.BoardUser, error) {
	foundUser := new(entities.BoardUser)
	err := repo.db.WithContext(ctx).
		Where(&entities.BoardUser{
			GoogleUserID: id,
		}).
		First(foundUser).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrBoardUserUserNotFound()
		}
		return nil, errors.NewAPIError(500, err)
	}

	return foundUser, nil
}

func (repo *BoardUserRepositoryInstance) FindByEmail(ctx context.Context, email string) (*entities.BoardUser, error) {
	foundUser := new(entities.BoardUser)
	err := repo.db.WithContext(ctx).
		Where(&entities.BoardUser{
			Email: email,
		}).
		First(foundUser).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrBoardUserIncorrectAuthen()
		}
		return nil, errors.NewAPIError(500, err)
	}

	return foundUser, nil
}

func (repo *BoardUserRepositoryInstance) Find(ctx context.Context) ([]entities.BoardUser, error) {
	foundUsers := make([]entities.BoardUser, 0)
	err := repo.db.WithContext(ctx).Find(&foundUsers).Error
	if err != nil {
		return nil, errors.NewAPIError(500, err)
	}

	return foundUsers, nil
}

func (repo *BoardUserRepositoryInstance) Create(ctx context.Context, newUser *entities.BoardUser) (*entities.BoardUser, error) {
	err := repo.db.WithContext(ctx).Create(newUser).Error
	if err != nil {
		return nil, errors.ErrBoardUserCreateFailed()
	}

	return newUser, nil
}
