package config

type GoogleAuthConfig struct {
	ClientID     string `env:"GAUTH_CLIENT_ID"`
	ClientSecret string `env:"GAUTH_CLIENT_SECRET"`
}

func NewGoogleAuthConfig(appConfig *AppConfig) *GoogleAuthConfig {
	return &appConfig.GAuth
}
