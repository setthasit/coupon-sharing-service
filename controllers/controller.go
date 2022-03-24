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
}

func NewRoute(
	authMiddlware *auth.AuthMiddleware,
	healthController HealthCheckController,
	boardUserController BoardUserController,
	boardController BoardController,
) *ControllerContainer {
	return &ControllerContainer{
		authMiddlware:       authMiddlware,
		healthController:    healthController,
		boardUserController: boardUserController,
		boardController:     boardController,
	}
}

func (cc *ControllerContainer) RegisterRoute(app *gin.Engine) {
	apiV1 := app.Group("/api/v1")

	apiV1.GET("/health", cc.healthController.HealthCheck)

	cc.registerBoardUserV1(apiV1)
	cc.registerBoardV1(apiV1)
}

func (cc *ControllerContainer) registerBoardUserV1(api *gin.RouterGroup) {
	userAPI := api.Group("/user")
	{
		// To be remove: for testing only
		userAPI.GET("", cc.boardUserController.GetUsers)
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
