package auth

import "github.com/google/wire"

var AuthProvider = wire.NewSet(
	NewGoogleOAuthClient,
)
