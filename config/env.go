package config

import (
	"log"

	env "github.com/Netflix/go-env"
	dotenv "github.com/joho/godotenv"
)

type AppConfig struct {
	Engine EngineConfig
	DB     DBConfig
	GAuth  GoogleAuthConfig
}

func NewAppConfig() *AppConfig {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	appConfig := new(AppConfig)
	_, err = env.UnmarshalFromEnviron(appConfig)
	if err != nil {
		log.Fatalf("Error Unmarshal config: %s", err)
	}

	return appConfig
}
