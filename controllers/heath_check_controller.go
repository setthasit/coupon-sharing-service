package controllers

import (
	"coupon-service/domains"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController interface {
	HealthCheck(c *gin.Context)
}

type HealthCheckControllerInstance struct {
	healthCheckSv domains.HealthCheckService
}

func NewHealthCheckController(healthCheckSv domains.HealthCheckService) *HealthCheckControllerInstance {
	return &HealthCheckControllerInstance{
		healthCheckSv: healthCheckSv,
	}
}

func (*HealthCheckControllerInstance) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
