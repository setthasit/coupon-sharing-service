package repositories

import (
	"context"
	"coupon-service/domains/entities"
	"coupon-service/infrastructure/errors"

	"gorm.io/gorm"
)

type CouponRepository interface {
	Find(ctx context.Context, boardID uint) ([]entities.Coupon, error)
	GetInfo(ctx context.Context, boardID uint, couponID uint) (*entities.Coupon, error)
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

func (repo *CouponRepositoryInstance) GetInfo(ctx context.Context, boardID uint, couponID uint) (*entities.Coupon, error) {
	foundCoupon := new(entities.Coupon)
	err := repo.db.
		WithContext(ctx).
		Where(&entities.Coupon{BoardID: boardID}).
		Preload("CouponUsageHistory", repo.db.Where(&entities.CouponUsageHistory{CouponID: couponID, BoardID: boardID})).
		Find(&foundCoupon).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrCouponNotFound()
		}
		return nil, errors.NewAPIError(500, err)
	}

	return foundCoupon, nil
}

func (repo *CouponRepositoryInstance) Create(ctx context.Context, newCoupon *entities.Coupon) (*entities.Coupon, error) {
	err := repo.db.WithContext(ctx).Create(newCoupon).Error
	if err != nil {
		return nil, errors.ErrCouponCannotCreate()
	}

	return newCoupon, nil
}

func (repo *CouponRepositoryInstance) CreateBatch(ctx context.Context, newCoupons []entities.Coupon) ([]entities.Coupon, error) {
	err := repo.db.WithContext(ctx).CreateInBatches(newCoupons, 100).Error
	if err != nil {
		return nil, errors.ErrCouponCannotCreate()
	}

	return newCoupons, nil
}
