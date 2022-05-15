package env

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port           int    `env:"PORT,unset" envDefault:"8000"`
	PostgresURL    string `env:"POSTGRES_CONNECTION_URL,unset"`
	AccessTokenKey string `env:"ACCESS_TOKEN_KEY,unset"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
