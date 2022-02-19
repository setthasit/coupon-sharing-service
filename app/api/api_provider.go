package api

import "github.com/google/wire"

var APIProvider = wire.NewSet(
	NewAPIContainer,
)
