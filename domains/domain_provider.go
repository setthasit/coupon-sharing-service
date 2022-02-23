package domains

import (
	"coupon-service/domains/services"

	"github.com/google/wire"
)

var ServiceProvider = wire.NewSet(
	services.NewHealthCheckService,
	services.NewBoardUserService,
)
