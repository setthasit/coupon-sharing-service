package controllers

import "github.com/gin-gonic/gin"

type ControllerContainer struct {
	healthController HealthCheckController
}

func NewRoute(
	healthController HealthCheckController,
) *ControllerContainer {
	return &ControllerContainer{
		healthController: healthController,
	}
}

func (cc *ControllerContainer) RegisterRoute(app *gin.Engine) {
	apiV1 := app.Group("/api/v1")

	apiV1.GET("/health", cc.healthController.HealthCheck)
}
