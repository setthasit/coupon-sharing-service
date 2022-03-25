package controllers

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CouponController interface {
	GetCouponByBoard(c *gin.Context)
	CreateCoupon(c *gin.Context)
	CreateBulkCoupons(c *gin.Context)
}

type CouponControllerInstance struct {
	couponSv    services.CouponService
	boardUserSv services.BoardUserService
}

func NewCouponController(
	couponSv services.CouponService,
	boardUserSv services.BoardUserService,
) CouponController {
	return &CouponControllerInstance{
		couponSv:    couponSv,
		boardUserSv: boardUserSv,
	}
}

func (ctrl *CouponControllerInstance) GetCouponByBoard(c *gin.Context) {
	// userInfo, err := getUserInfoFromGinContext(c, ctrl.boardUserSv)
	// if err != nil {
	// 	handleError(c, err)
	// 	return
	// }

	// responseItemHttp(c, http.StatusOK, coupon)
}

func (ctrl *CouponControllerInstance) CreateCoupon(c *gin.Context) {
	userInfo, err := getUserInfoFromGinContext(c, ctrl.boardUserSv)
	if err != nil {
		handleError(c, err)
		return
	}

	newCouponReq := new(entities.CouponCreateNew)
	err = c.BindJSON(newCouponReq)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, err.Error())
		return
	}

	coupon, err := ctrl.couponSv.CreateNewCoupon(c, newCouponReq, userInfo.ID)
	if err != nil {
		handleError(c, err)
		return
	}

	responseItemHttp(c, http.StatusOK, coupon)
}

func (ctrl *CouponControllerInstance) CreateBulkCoupons(c *gin.Context) {
	userInfo, err := getUserInfoFromGinContext(c, ctrl.boardUserSv)
	if err != nil {
		handleError(c, err)
		return
	}

	newCouponReq := new(entities.CouponBulkCreateNew)
	err = c.BindJSON(newCouponReq)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, err.Error())
		return
	}

	coupons, err := ctrl.couponSv.BulkCreateCoupons(c, newCouponReq, userInfo.ID)
	if err != nil {
		handleError(c, err)
		return
	}

	responseListHttp(c, http.StatusOK, coupons, len(coupons))
}
