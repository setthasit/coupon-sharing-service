package repositories

import (
	"context"
	"coupon-service/domains/entities"
	"coupon-service/infrastructure/errors"

	"gorm.io/gorm"
)

type CouponRepository interface {
	Find(ctx context.Context, boardID uint) ([]entities.Coupon, error)
	Create(ctx context.Context, newCoupon *entities.Coupon) (*entities.Coupon, error)
	CreateBatch(ctx context.Context, newCoupons []entities.Coupon) ([]entities.Coupon, error)
}

type CouponRepositoryInstance struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) CouponRepository {
	return &CouponRepositoryInstance{
		db: db,
	}
}

func (repo *CouponRepositoryInstance) Find(ctx context.Context, boardID uint) ([]entities.Coupon, error) {
	foundCoupons := make([]entities.Coupon, 0)
	err := repo.db.
		WithContext(ctx).
		Where(&entities.Coupon{BoardID: boardID}).
		Find(&foundCoupons).
		Error
	if err != nil {
		return nil, errors.NewAPIError(500, err)
	}

	return foundCoupons, nil
}

func (repo *CouponRepositoryInstance) Create(ctx context.Context, newCoupon *entities.Coupon) (*entities.Coupon, error) {
	err := repo.db.WithContext(ctx).Create(newCoupon).Error
	if err != nil {
		return nil, errors.ErrAuthInvalidToken()
	}

	return newCoupon, nil
}

func (repo *CouponRepositoryInstance) CreateBatch(ctx context.Context, newCoupons []entities.Coupon) ([]entities.Coupon, error) {
	err := repo.db.WithContext(ctx).CreateInBatches(newCoupons, 100).Error
	if err != nil {
		return nil, errors.ErrAuthInvalidToken()
	}

	return newCoupons, nil
}
