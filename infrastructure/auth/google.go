package auth

import (
	"context"
	"coupon-service/config"

	"google.golang.org/api/idtoken"
	"google.golang.org/api/oauth2/v2"
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

func (gOAuth *GoogleOAuth) GetUserData(ctx context.Context, tokenID string) (*oauth2.Userinfo, error) {
	payload, err := idtoken.Validate(ctx, tokenID, gOAuth.ClientID)
	if err != nil {
		return nil, err
	}

	info, err := parseUserInfoFromClaim(payload.Claims)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func parseUserInfoFromClaim(claim map[string]interface{}) (*oauth2.Userinfo, error) {
	verifiedEmail, _ := claim["email_verified"].(*bool)

	return &oauth2.Userinfo{
		Id:            claim["sub"].(string),
		Email:         claim["email"].(string),
		Name:          claim["name"].(string),
		GivenName:     claim["given_name"].(string),
		FamilyName:    claim["family_name"].(string),
		Picture:       claim["picture"].(string),
		Locale:        claim["locale"].(string),
		VerifiedEmail: verifiedEmail,
	}, nil
}
