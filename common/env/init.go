package env

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port           int    `env:"PORT,unset" envDefault:"8000"`
	BaseUrl        string `env:"BASE_URL,unset"`
	PostgresURL    string `env:"POSTGRES_CONNECTION_URL,unset"`
	MigrationPath  string `env:"POSTGRES_MIGRATION_PATH,unset"`
	InvoicePath    string `env:"INVOICE_CONFIG_PATH,unset"`
	AccessTokenKey string `env:"ACCESS_TOKEN_KEY,unset"`
	MailEmail      string `env:"MAIL_EMAIL,unset"`
	MailPassword   string `env:"MAIL_PASSWORD,unset"`
	MailSmtpHost   string `env:"MAIL_SMTP_HOST,unset"`
	GinMode        string `env:"GIN_MODE,unset" envDefault:"debug"`
	MailSmtpPort   int    `env:"MAIL_SMTP_PORT,unset"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
