package domains

import "github.com/google/wire"

var ServiceProvider = wire.NewSet(
	NewHealthCheckService,
)
