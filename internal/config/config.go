package config

import "luizalabs-technical-test/pkg/env"

var (
	Server        serverConfig
	ExternalAPI   externalAPI
	GeneralConfig generalConfig
)

func LoadEnv() {
	const TAG_NAME = "env"
	env.LoadStructWithEnvVars(TAG_NAME, &Server, &ExternalAPI, &GeneralConfig)
}

type generalConfig struct {
	SECRET_AUTH_TOKEN_KEY string `env:"SECRET_AUTH_TOKEN_KEY"`
}

type serverConfig struct {
	PORT string `env:"SERVER_PORT"`
	HOST string `env:"SERVER_HOST"`
}

type externalAPI struct {
	URL string `env:"EXTERNAL_API"`
}
