package config

import (
	"fmt"
	"luizalabs-technical-test/pkg/env"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test loading environment variables into configuration structures.
func TestInit(t *testing.T) {
	// ARRANGE & ACT
	envVars := map[string]string{
		"PG_HOST":               "localhost",
		"PG_PORT":               "5432",
		"PG_USER":               "user",
		"PG_PASSWORD":           "password",
		"PG_DATABASE":           "testdb",
		"SECRET_AUTH_TOKEN_KEY": "test-secret",
		"SERVER_PORT":           "8080",
		"SERVER_HOST":           "localhost",
	}

	for key, value := range envVars {
		err := os.Setenv(key, value)
		assert.NoError(t, err, "failed to set environment variable")
	}
	env.LoadStructWithEnvVars("env", &ServerConfig, &GeneralConfig, &PostgresConfig)

	// ASSERT
	assert.Equal(t, envVars["PG_HOST"], PostgresConfig.Host)
	assert.Equal(t, envVars["PG_PORT"], PostgresConfig.Port)
	assert.Equal(t, envVars["PG_USER"], PostgresConfig.User)
	assert.Equal(t, envVars["PG_PASSWORD"], PostgresConfig.Password)
	assert.Equal(t, envVars["PG_DATABASE"], PostgresConfig.Database)
	assert.Equal(t, envVars["SECRET_AUTH_TOKEN_KEY"], GeneralConfig.SecretAuthTokenKey)
	assert.Equal(t, envVars["SERVER_PORT"], ServerConfig.Port)
	assert.Equal(t, envVars["SERVER_HOST"], ServerConfig.Host)
}

// Test ToPostgresDSN method in order to correct generate connection string.
func TestToPostgresDSN(t *testing.T) {
	// ARRANGE
	p := postgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "user",
		Password: "password",
		Database: "testdb",
	}

	// ACT & ASSERT
	expectedDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		p.Host, p.User, p.Password, p.Database, p.Port)

	assert.Equal(t, expectedDSN, p.ToPostgresDSN())
}
