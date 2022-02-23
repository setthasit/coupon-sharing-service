package controllers

import (
	"coupon-service/domains/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController interface {
	HealthCheck(c *gin.Context)
}

type HealthCheckControllerInstance struct {
	healthCheckSv services.HealthCheckService
}

func NewHealthCheckController(healthCheckSv services.HealthCheckService) HealthCheckController {
	return &HealthCheckControllerInstance{
		healthCheckSv: healthCheckSv,
	}
}

func (*HealthCheckControllerInstance) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
