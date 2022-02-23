package config

import "github.com/google/wire"

var ConfigProvider = wire.NewSet(
	NewAppConfig,
	NewDBConfig,
)
