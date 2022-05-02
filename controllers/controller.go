package controllers

import (
	"coupon-service/infrastructure/auth"

	"github.com/gin-gonic/gin"
)

type ControllerContainer struct {
	authMiddlware       *auth.AuthMiddleware
	healthController    HealthCheckController
	boardUserController BoardUserController
	boardController     BoardController
	couponController    CouponController
}

func NewRoute(
	authMiddlware *auth.AuthMiddleware,
	healthController HealthCheckController,
	boardUserController BoardUserController,
	boardController BoardController,
	couponController CouponController,
) *ControllerContainer {
	return &ControllerContainer{
		authMiddlware:       authMiddlware,
		healthController:    healthController,
		boardUserController: boardUserController,
		boardController:     boardController,
		couponController:    couponController,
	}
}

func (cc *ControllerContainer) RegisterRoute(app *gin.Engine) {
	apiV1 := app.Group("/api/v1")

	apiV1.GET("/health", cc.healthController.HealthCheck)

	cc.registerBoardUserV1(apiV1)
	cc.registerBoardV1(apiV1)
	cc.registerCouponV1(apiV1)
}

func (cc *ControllerContainer) registerBoardUserV1(api *gin.RouterGroup) {
	userAPI := api.Group("/user")
	{
		userAPI.POST("/signin/google", cc.boardUserController.SignInGoogle)

		authUserAPI := userAPI.Use(cc.authMiddlware.Register())
		{
			authUserAPI.GET("/me", cc.boardUserController.GetSelfInfo)
		}
	}
}

func (cc *ControllerContainer) registerBoardV1(api *gin.RouterGroup) {
	boardAPI := api.Group("/board")
	{
		authUserAPI := boardAPI.Use(cc.authMiddlware.Register())
		{
			authUserAPI.GET("", cc.boardController.GetBoardByUser)
			authUserAPI.POST("", cc.boardController.CreateNewBoard)
		}
	}
}

func (cc *ControllerContainer) registerCouponV1(api *gin.RouterGroup) {
	couponAPI := api.Group("/coupon")
	{
		authUserAPI := couponAPI.Use(cc.authMiddlware.Register())
		{
			authUserAPI.GET("/:board_id", cc.couponController.GetCouponByBoard)
			authUserAPI.GET("/:board_id/:coupon_id", cc.couponController.GetCouponInfo)
			authUserAPI.POST("", cc.couponController.CreateCoupon)
			authUserAPI.POST("/bulk", cc.couponController.CreateBulkCoupons)

			authUserAPI.POST("/copy", cc.couponController.Copy)
			authUserAPI.POST("/use", cc.couponController.Use)
			authUserAPI.POST("/cancel", cc.couponController.Cancel)
		}
	}
}
