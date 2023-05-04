package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort       string
	Postgres       PostgresConfig
	AuthSecretKey  string
	LogLevel       string
	SiginKey       string
	AuthConfigPath string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load(path string) Config {
	godotenv.Load(path + "/.env") // load .env file if it exists

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		HttpPort: conf.GetString("HTTP_PORT"),
		Postgres: PostgresConfig{
			Host:     conf.GetString("POSTGRES_HOST"),
			Port:     conf.GetString("POSTGRES_PORT"),
			User:     conf.GetString("POSTGRES_USER"),
			Password: conf.GetString("POSTGRES_PASSWORD"),
			Database: conf.GetString("POSTGRES_DATABASE"),
		},
		AuthSecretKey:  conf.GetString("AUTH_SECRET_KEY"),
		LogLevel:       conf.GetString("LOG_LEVEL"),
		SiginKey:       conf.GetString("SIGNING_KEY"),
		AuthConfigPath: conf.GetString("CASBIN_CONFIG_PATH"),
	}

	return cfg
}
