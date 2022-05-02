package repositories

import (
	"context"
	"coupon-service/domains/entities"
	"coupon-service/infrastructure/errors"

	"gorm.io/gorm"
)

type CouponUsageHistoryRepository interface {
	Copy(ctx context.Context, coupon *entities.CouponUsageHistory) error
	Use(ctx context.Context, coupon *entities.CouponUsageHistory) error
	Cancel(ctx context.Context, coupon *entities.CouponUsageHistory) error
}

type CouponUsageHistoryRepositoryInstance struct {
	db *gorm.DB
}

func NewCouponUsageHistoryRepository(db *gorm.DB) CouponUsageHistoryRepository {
	return &CouponUsageHistoryRepositoryInstance{
		db: db,
	}
}

func (repo *CouponUsageHistoryRepositoryInstance) Copy(ctx context.Context, coupon *entities.CouponUsageHistory) error {
	err := repo.db.WithContext(ctx).Create(coupon).Error
	if err != nil {
		return errors.ErrCouponHistoyUsageFailedCopyAction()
	}
	return nil
}

func (repo *CouponUsageHistoryRepositoryInstance) Use(ctx context.Context, coupon *entities.CouponUsageHistory) error {
	err := repo.db.WithContext(ctx).Create(coupon).Error
	if err != nil {
		return errors.ErrCouponHistoyUsageFailedUseAction()
	}
	return nil
}

func (repo *CouponUsageHistoryRepositoryInstance) Cancel(ctx context.Context, coupon *entities.CouponUsageHistory) error {
	err := repo.db.WithContext(ctx).Create(coupon).Error
	if err != nil {
		return errors.ErrCouponHistoyUsageFailedCancelAction()
	}
	return nil
}
