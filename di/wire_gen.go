// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"coupon-service/app/api"
	"coupon-service/controllers"
	"coupon-service/domains"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeApp() *api.APIContainer {
	healthCheckServiceInstance := domains.NewHealthCheckService()
	healthCheckControllerInstance := controllers.NewHealthCheckController(healthCheckServiceInstance)
	controllerContainer := controllers.NewRoute(healthCheckControllerInstance)
	apiContainer := api.NewAPIContainer(controllerContainer)
	return apiContainer
}

// wire.go:

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

var BindingSet = wire.NewSet(api.APIProvider, wire.Struct(new(Config), "*"), wire.Struct(new(Repositories), "*"), domains.ServiceProvider, wire.Struct(new(Services), "*"), controllers.ControllerProvider, wire.Struct(new(Controllers), "*"), wire.Bind(new(domains.HealthCheckService), new(*domains.HealthCheckServiceInstance)), wire.Bind(new(controllers.HealthCheckController), new(*controllers.HealthCheckControllerInstance)))
