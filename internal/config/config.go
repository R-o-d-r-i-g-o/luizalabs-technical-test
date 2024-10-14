package config

import (
	"fmt"
	"luizalabs-technical-test/pkg/env"

	"github.com/joho/godotenv"
)

// Variables that store server, external API, and general configuration settings.
var (
	ServerConfig   serverConfig
	GeneralConfig  generalConfig
	PostgresConfig postgresConfig
)

// init loads environment variables into the configuration structures using "env" tags.
func init() {
	const tagName = "env"

	godotenv.Load(".env")
	env.LoadStructWithEnvVars(tagName, &ServerConfig, &GeneralConfig, &PostgresConfig)
}

// Structure to load database configurations (connection string).
type postgresConfig struct {
	Host     string `env:"PG_HOST"`
	Port     string `env:"PG_PORT"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Database string `env:"PG_DATABASE"`
}

// Structure to load general configurations (e.g., authentication key).
type generalConfig struct {
	SecretAuthTokenKey string `env:"SECRET_AUTH_TOKEN_KEY"`
}

// Structure to load server configurations (port and host).
type serverConfig struct {
	Port string `env:"SERVER_PORT"`
	Host string `env:"SERVER_HOST"`
}

// ToPostgresDSN fromats provided data into postgres db dsn.
func (p *postgresConfig) ToPostgresDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		p.Host, p.User, p.Password, p.Database, p.Port,
	)
}
