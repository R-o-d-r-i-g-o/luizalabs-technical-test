package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// PasswordHasherSuite is a test suite for password hashing functionalities.
type PasswordHasherSuite struct {
	suite.Suite
	hasher PasswordHasher
}

// SetupSuite sets up the test suite.
func (s *PasswordHasherSuite) SetupSuite() {
	s.hasher = &passwordHasher{}
}

// TestHashPassword tests the Hash function of the PasswordHasher.
func (s *PasswordHasherSuite) TestHashPassword() {
	// ARRANGE & ACT
	password := "my_secret_password"
	hashedPassword, err := s.hasher.HashPassword(password)

	// ASSERT
	assert.NoError(s.T(), err, "Expected no error while hashing the password")
	assert.NotEmpty(s.T(), hashedPassword, "Hashed password should not be empty")
	assert.NotEqual(s.T(), password, hashedPassword, "Hashed password should not equal the plain password")
}

// TestCheckPasswordHash tests the Check function of the PasswordHasher.
func (s *PasswordHasherSuite) TestCheckPasswordHash() {
	// ARRANGE & ACT
	password := "my_secret_password"
	hashedPassword, err := s.hasher.HashPassword(password)
	if err != nil {
		s.T().Fatalf("Failed to hash password: %v", err)
	}

	// ASSERT
	assert.True(s.T(), s.hasher.CheckPasswordHash(password, hashedPassword), "Expected password to match the hash")
	assert.False(s.T(), s.hasher.CheckPasswordHash("wrong_password", hashedPassword), "Expected wrong password not to match the hash")
}

// TestPasswordHasherSuite runs the test suite.
func TestPasswordHasherSuite(t *testing.T) {
	suite.Run(t, new(PasswordHasherSuite))
}
