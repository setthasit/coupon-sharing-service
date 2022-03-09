package config

type DBConfig struct {
	Host         string `env:"DB_HOST"`
	Username     string `env:"DB_USERNAME"`
	Password     string `env:"DB_PASSWORD"`
	Port         string `env:"DB_PORT"`
	DatabaseName string `env:"DB_DATABASE_NAME"`
	IsMigration  bool   `env:"DB_MIGRATION,default=false"`
}

func NewDBConfig(appConfig *AppConfig) *DBConfig {
	return &appConfig.DB
}
