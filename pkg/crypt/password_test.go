package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "my_secret_password"

	// Test hashing the password
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	// Test that the hashed password is not the same as the plain password
	assert.NotEqual(t, password, hashedPassword)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "my_secret_password"
	hashedPassword, _ := HashPassword(password)

	// Test correct password
	assert.True(t, CheckPasswordHash(password, hashedPassword))

	// Test incorrect password
	assert.False(t, CheckPasswordHash("wrong_password", hashedPassword))
}
