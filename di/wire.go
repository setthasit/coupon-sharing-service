//go:build wireinject
// +build wireinject

package di

import (
	"coupon-service/app/api"
	"coupon-service/controllers"
	"coupon-service/domains"

	"github.com/google/wire"
)

func InitializeApp() *api.APIContainer {
	wire.Build(BindingSet)

	return &api.APIContainer{}
}

// Config
type Config struct{}

// Controllers
type Controllers struct {
	HealthCheckController controllers.HealthCheckController
}

// Services
type Services struct {
	HealthCheckService domains.HealthCheckService
}

// Repositories
type Repositories struct{}

var BindingSet = wire.NewSet(
	api.APIProvider,

	wire.Struct(new(Config), "*"),
	wire.Struct(new(Repositories), "*"),

	domains.ServiceProvider,
	wire.Struct(new(Services), "*"),

	controllers.ControllerProvider,
	wire.Struct(new(Controllers), "*"),

	// Services
	wire.Bind(new(domains.HealthCheckService), new(*domains.HealthCheckServiceInstance)),
	// Controllers
	wire.Bind(new(controllers.HealthCheckController), new(*controllers.HealthCheckControllerInstance)),
)
