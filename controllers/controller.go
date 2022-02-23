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

	apiV1.POST("/user/register", cc.boardUserController.CreateNewUser)
}
