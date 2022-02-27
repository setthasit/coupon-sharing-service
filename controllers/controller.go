package controllers

import "github.com/gin-gonic/gin"

type ControllerContainer struct {
	healthController    HealthCheckController
	boardUserController BoardUserController
}

func NewRoute(
	healthController HealthCheckController,
	boardUserController BoardUserController,
) *ControllerContainer {
	return &ControllerContainer{
		healthController:    healthController,
		boardUserController: boardUserController,
	}
}

func (cc *ControllerContainer) RegisterRoute(app *gin.Engine) {
	apiV1 := app.Group("/api/v1")

	apiV1.GET("/health", cc.healthController.HealthCheck)

	userAPI := apiV1.Group("/user")
	{
		// To be remove: for testing only
		userAPI.GET("", cc.boardUserController.GetUsers)
		userAPI.POST("/register", cc.boardUserController.CreateNewUser)

		userAPI.POST("/signin", cc.boardUserController.SignIn)
	}
}
