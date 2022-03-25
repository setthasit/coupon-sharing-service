package domains

import (
	"coupon-service/domains/repositories"
	"coupon-service/domains/services"

	"github.com/google/wire"
)

var ServiceProvider = wire.NewSet(
	// Sertvices
	services.NewHealthCheckService,
	services.NewBoardUserService,
	services.NewBoardService,
	services.NewCouponService,

	// Repositories
	repositories.NewBoardUserRepository,
	repositories.NewBoardRepository,
	repositories.NewBoardMemberRepository,
	repositories.NewCouponRepository,
)
