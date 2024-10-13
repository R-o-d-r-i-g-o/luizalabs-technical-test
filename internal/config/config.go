package config

import "luizalabs-technical-test/pkg/env"

// Variables that store server, external API, and general configuration settings
var (
	ServerEnv        serverConfig
	ExternalAPIEnv   externalAPI
	GeneralConfigEnv generalConfig
	DatabaseEnv      databaseConfig
)

// LoadEnv loads environment variables into the configuration structures using "env" tags
func LoadEnv() {
	const tagName = "env"
	env.LoadStructWithEnvVars(tagName, &ServerEnv, &ExternalAPIEnv, &GeneralConfigEnv, &DatabaseEnv)
}

// Structure to load database configurations (connection string)
type databaseConfig struct {
	ConnectionString string `env:"DATABASE_CONNECTION_STRING"`
}

// Structure to load general configurations (e.g., authentication key)
type generalConfig struct {
	SecretAuthTokenKey string `env:"SECRET_AUTH_TOKEN_KEY"`
}

// Structure to load server configurations (port and host)
type serverConfig struct {
	Port string `env:"SERVER_PORT"`
	Host string `env:"SERVER_HOST"`
}

// Structure to load the external API URL
type externalAPI struct {
	URL string `env:"EXTERNAL_API"`
}
