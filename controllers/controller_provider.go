package controllers

import "github.com/google/wire"

var ControllerProvider = wire.NewSet(
	NewRoute,
	NewHealthCheckController,
	NewBoardUserController,
	NewBoardController,
	NewCouponController,
)
