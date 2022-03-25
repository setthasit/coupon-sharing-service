package services

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/repositories"

	"github.com/gin-gonic/gin"
)

type CouponService interface {
	CreateNewCoupon(ctx *gin.Context, newCoupon *entities.CouponCreateNew, boardUserID uint) (*entities.Coupon, error)
	BulkCreateCoupons(ctx *gin.Context, newCoupons *entities.CouponBulkCreateNew, boardUserID uint) ([]entities.Coupon, error)
}

type CouponServiceInstance struct {
	couponRepo repositories.CouponRepository
}

func NewCouponService(
	couponRepo repositories.CouponRepository,
) CouponService {
	return &CouponServiceInstance{
		couponRepo: couponRepo,
	}
}

func (sv *CouponServiceInstance) CreateNewCoupon(ctx *gin.Context, newCoupon *entities.CouponCreateNew, boardUserID uint) (*entities.Coupon, error) {
	coupon, err := sv.couponRepo.Create(ctx, newCoupon.ToCouponWithAudit(boardUserID))
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (sv *CouponServiceInstance) BulkCreateCoupons(ctx *gin.Context, newCoupons *entities.CouponBulkCreateNew, boardUserID uint) ([]entities.Coupon, error) {
	coupons, err := sv.couponRepo.CreateBatch(ctx, newCoupons.ToCouponsWithAudit(boardUserID))
	if err != nil {
		return nil, err
	}

	return coupons, nil
}
