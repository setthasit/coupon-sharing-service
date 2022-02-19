package api

import (
	"coupon-service/controllers"

	"github.com/gin-gonic/gin"
)

type APIContainer struct {
	Engine *gin.Engine
	Ctrl   *controllers.ControllerContainer
}

func NewAPIContainer(
	ctrl *controllers.ControllerContainer,
) *APIContainer {
	return &APIContainer{
		Engine: gin.New(),
		Ctrl:   ctrl,
	}
}
