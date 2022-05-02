package services

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/repositories"

	"github.com/gin-gonic/gin"
)

type CouponService interface {
	GetInfo(ctx *gin.Context, boardID uint, couponID uint, boardUserID uint) (*entities.Coupon, error)
	FindByBoardID(ctx *gin.Context, boardID uint, boardUserID uint) ([]entities.Coupon, error)
	CreateNewCoupon(ctx *gin.Context, newCoupon *entities.CouponCreateNew, boardUserID uint) (*entities.Coupon, error)
	BulkCreateCoupons(ctx *gin.Context, newCoupons *entities.CouponBulkCreateNew, boardUserID uint) ([]entities.Coupon, error)
	Copy(ctx *gin.Context, req *entities.CouponCopyRequest) error
	Use(ctx *gin.Context, req *entities.CouponUseRequest) error
	Cancel(ctx *gin.Context, req *entities.CouponCancelRequest) error
}

type CouponServiceInstance struct {
	couponRepo      repositories.CouponRepository
	couponUsageRepo repositories.CouponUsageHistoryRepository
	boardMemberRepo repositories.BoardMemberRepository
}

func NewCouponService(
	couponRepo repositories.CouponRepository,
	couponUsageRepo repositories.CouponUsageHistoryRepository,
	boardMemberRepo repositories.BoardMemberRepository,
) CouponService {
	return &CouponServiceInstance{
		couponRepo:      couponRepo,
		couponUsageRepo: couponUsageRepo,
		boardMemberRepo: boardMemberRepo,
	}
}

func (sv *CouponServiceInstance) FindByBoardID(ctx *gin.Context, boardID uint, boardUserID uint) ([]entities.Coupon, error) {
	err := sv.boardMemberRepo.ValidateUserInBoard(ctx, boardID, boardUserID)
	if err != nil {
		return nil, err
	}

	coupon, err := sv.couponRepo.Find(ctx, boardID)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (sv *CouponServiceInstance) GetInfo(ctx *gin.Context, boardID uint, couponID uint, boardUserID uint) (*entities.Coupon, error) {
	err := sv.boardMemberRepo.ValidateUserInBoard(ctx, boardID, boardUserID)
	if err != nil {
		return nil, err
	}

	coupon, err := sv.couponRepo.GetInfo(ctx, boardID, couponID)
	if err != nil {
		return nil, err
	}
	return coupon, nil
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

func (sv *CouponServiceInstance) Copy(ctx *gin.Context, req *entities.CouponCopyRequest) error {
	err := sv.boardMemberRepo.ValidateUserInBoard(ctx, req.BoardID, req.BoardUserID)
	if err != nil {
		return err
	}

	return sv.couponUsageRepo.Copy(ctx, req.ToCouponUsageHistory())
}

func (sv *CouponServiceInstance) Use(ctx *gin.Context, req *entities.CouponUseRequest) error {
	err := sv.boardMemberRepo.ValidateUserInBoard(ctx, req.BoardID, req.BoardUserID)
	if err != nil {
		return err
	}

	return sv.couponUsageRepo.Copy(ctx, req.ToCouponUsageHistory())
}

func (sv *CouponServiceInstance) Cancel(ctx *gin.Context, req *entities.CouponCancelRequest) error {
	err := sv.boardMemberRepo.ValidateUserInBoard(ctx, req.BoardID, req.BoardUserID)
	if err != nil {
		return err
	}

	return sv.couponUsageRepo.Copy(ctx, req.ToCouponUsageHistory())
}
