package api

import (
	"coupon-service/config"
	"coupon-service/controllers"
	"coupon-service/infrastructure/security"

	"github.com/gin-gonic/gin"
)

type APIContainer struct {
	Engine *gin.Engine
	Ctrl   *controllers.ControllerContainer
}

func NewAPIContainer(
	cfg *config.EngineConfig,
	ctrl *controllers.ControllerContainer,
) *APIContainer {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(security.CORSMiddleware(cfg))

	if cfg.TrustCloudflare {
		app.TrustedPlatform = gin.PlatformCloudflare
	}

	return &APIContainer{
		Engine: app,
		Ctrl:   ctrl,
	}
}
