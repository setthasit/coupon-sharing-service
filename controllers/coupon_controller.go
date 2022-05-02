package controllers

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CouponController interface {
	GetCouponByBoard(c *gin.Context)
	GetCouponInfo(c *gin.Context)
	CreateCoupon(c *gin.Context)
	CreateBulkCoupons(c *gin.Context)
	Copy(c *gin.Context)
	Use(c *gin.Context)
	Cancel(c *gin.Context)
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
	userInfo, err := getUserInfoFromGinContext(c, ctrl.boardUserSv)
	if err != nil {
		handleError(c, err)
		return
	}

	paramBoardID := c.Param("board_id")
	if len(paramBoardID) < 1 {
		responsMessageHttp(c, http.StatusBadRequest, "board not found")
		return
	}

	boardID, err := strconv.Atoi(paramBoardID)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, "invalid board id")
		return
	}

	coupons, err := ctrl.couponSv.FindByBoardID(c, uint(boardID), userInfo.ID)
	if err != nil {
		handleError(c, err)
		return
	}

	responseListHttp(c, http.StatusOK, coupons, len(coupons))
}

func (ctrl *CouponControllerInstance) GetCouponInfo(c *gin.Context) {
	userInfo, err := getUserInfoFromGinContext(c, ctrl.boardUserSv)
	if err != nil {
		handleError(c, err)
		return
	}

	paramBoardID := c.Param("board_id")
	if len(paramBoardID) < 1 {
		responsMessageHttp(c, http.StatusBadRequest, "board not found")
		return
	}

	boardID, err := strconv.Atoi(paramBoardID)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, "invalid board id")
		return
	}

	paramCouponID := c.Param("coupon_id")
	if len(paramCouponID) < 1 {
		responsMessageHttp(c, http.StatusBadRequest, "board not found")
		return
	}

	couponID, err := strconv.Atoi(paramCouponID)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, "invalid board id")
		return
	}

	coupon, err := ctrl.couponSv.GetInfo(c, uint(boardID), uint(couponID), userInfo.ID)
	if err != nil {
		handleError(c, err)
		return
	}

	responseItemHttp(c, http.StatusOK, coupon)
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

func (ctrl *CouponControllerInstance) Copy(c *gin.Context) {
	userInfo, err := getUserInfoFromGinContext(c, ctrl.boardUserSv)
	if err != nil {
		handleError(c, err)
		return
	}

	copyRequest := new(entities.CouponCopyRequest)
	err = c.BindJSON(copyRequest)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, err.Error())
		return
	}
	copyRequest.BoardUserID = userInfo.ID

	err = ctrl.couponSv.Copy(c, copyRequest)
	if err != nil {
		handleError(c, err)
		return
	}

	responseEmptyHttp(c, http.StatusOK)
}

func (ctrl *CouponControllerInstance) Use(c *gin.Context) {
	userInfo, err := getUserInfoFromGinContext(c, ctrl.boardUserSv)
	if err != nil {
		handleError(c, err)
		return
	}

	useRequest := new(entities.CouponUseRequest)
	err = c.BindJSON(useRequest)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, err.Error())
		return
	}
	useRequest.BoardUserID = userInfo.ID

	err = ctrl.couponSv.Use(c, useRequest)
	if err != nil {
		handleError(c, err)
		return
	}

	responseEmptyHttp(c, http.StatusOK)
}

func (ctrl *CouponControllerInstance) Cancel(c *gin.Context) {
	userInfo, err := getUserInfoFromGinContext(c, ctrl.boardUserSv)
	if err != nil {
		handleError(c, err)
		return
	}

	cancelRequest := new(entities.CouponCancelRequest)
	err = c.BindJSON(cancelRequest)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, err.Error())
		return
	}
	cancelRequest.BoardUserID = userInfo.ID

	err = ctrl.couponSv.Cancel(c, cancelRequest)
	if err != nil {
		handleError(c, err)
		return
	}

	responseEmptyHttp(c, http.StatusOK)
}
