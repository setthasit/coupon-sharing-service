package persistence

import "github.com/google/wire"

var PersistenceProvider = wire.NewSet(
	NewDbConn,
	NewBoardUserRepository,
)
