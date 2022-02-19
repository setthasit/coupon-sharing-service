package main

import (
	"coupon-service/di"
)

func main() {
	app := di.InitializeApp()

	app.Ctrl.RegisterRoute(app.Engine)

	app.Engine.Run(":9999")
}
