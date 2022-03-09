package auth

import (
	"context"
	"coupon-service/config"

	"google.golang.org/api/idtoken"
)

type GoogleOAuth struct {
	ClientID     string
	ClientSecret string
}

func NewGoogleOAuthClient(gAuthConfig *config.GoogleAuthConfig) *GoogleOAuth {
	return &GoogleOAuth{
		ClientID:     gAuthConfig.ClientID,
		ClientSecret: gAuthConfig.ClientSecret,
	}
}

func (gOAuth *GoogleOAuth) GetUserData(ctx context.Context, tokenID string) (interface{}, error) {
	payload, err := idtoken.Validate(ctx, tokenID, gOAuth.ClientID)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
