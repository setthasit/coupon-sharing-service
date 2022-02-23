//go:build wireinject
// +build wireinject

package di

import (
	"coupon-service/app/api"
	"coupon-service/config"
	"coupon-service/controllers"
	"coupon-service/domains"
	"coupon-service/infrastructure/persistence"

	"github.com/google/wire"
)

func InitializeApp() *api.APIContainer {
	wire.Build(BindingSet)

	return &api.APIContainer{}
}

var BindingSet = wire.NewSet(
	api.APIProvider,
	config.ConfigProvider,
	persistence.PersistenceProvider,
	domains.ServiceProvider,
	controllers.ControllerProvider,
)
