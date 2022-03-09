package config

type EngineConfig struct {
	Port string `env:"ENGINE_SERVE_PORT,default=9999"`

	TrustCloudflare bool   `env:"ENGINE_TRUST_CLOUDFLARE,default=false"`
	AllowCors       string `env:"ENGINE_ALLOW_CORS,default=http://localhost:3000"`
}

func NewEngineConfig(appConfig *AppConfig) *EngineConfig {
	return &appConfig.Engine
}
