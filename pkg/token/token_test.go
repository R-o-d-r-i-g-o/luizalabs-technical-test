package token_test

import (
	"luizalabs-technical-test/pkg/token"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TokenTestSuite struct {
	suite.Suite
}

func (suite *TokenTestSuite) TestCreateTokenAndValidateToken() {
	// Set up test data
	testClaims := token.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "test-issuer",
		},
		CustomKeys: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
		},
	}

	// Create token
	secretKey := "secret_key"
	tokenString, err := token.CreateToken(secretKey, testClaims)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), tokenString)

	// Validate token
	claims, err := token.ValidateToken(secretKey, tokenString)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), claims)

	// Verify token claims
	assert.Equal(suite.T(), testClaims.StandardClaims.ExpiresAt, claims.ExpiresAt)
	assert.Equal(suite.T(), testClaims.StandardClaims.Issuer, claims.Issuer)
	assert.Equal(suite.T(), testClaims.CustomKeys["key1"], claims.CustomKeys["key1"])
	assert.EqualValues(suite.T(), testClaims.CustomKeys["key2"], claims.CustomKeys["key2"])
}

func (suite *TokenTestSuite) TestValidateInvalidToken() {
	// Attempt to validate an invalid token
	secretKey := "secret_key"
	invalidTokenString := "invalid_token_string"
	claims, err := token.ValidateToken(secretKey, invalidTokenString)

	// Validate the error and claims
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), claims)

	// Ensure that the error is of the expected type
	assert.Contains(suite.T(), err.Error(), "token contains an invalid number of segments")
}

func (suite *TokenTestSuite) TestValidateTokenWithUnexpectedSigningMethod() {
	// Create a token with an unexpected signing method
	testClaims := token.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "test-issuer",
		},
		CustomKeys: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
		},
	}

	firstSecretKey := "first_secret_key"
	tokenString, err := token.CreateToken(firstSecretKey, testClaims)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), tokenString)

	// Attempt to validate the token with an unexpected signing method
	secondSecretKey := "second_secret_key"
	claims, err := token.ValidateToken(secondSecretKey, tokenString)

	// Validate the error and claims
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), claims)

	// Ensure that the error is of the expected type
	assert.Contains(suite.T(), err.Error(), "signature is invalid")
}

func (suite *TokenTestSuite) TestValidateTokenWithExpiredToken() {
	// Set up test data
	testClaims := token.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Hour).Unix(), // Expired token
			Issuer:    "test-issuer",
		},
		CustomKeys: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
		},
	}

	// Create an expired token
	secretKey := "secret_key"
	expiredTokenString, err := token.CreateToken(secretKey, testClaims)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), expiredTokenString)

	// Attempt to validate the expired token
	claims, err := token.ValidateToken(secretKey, expiredTokenString)

	// Validate the error and claims
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), claims)

	// Ensure that the error is of the expected type
	assert.Contains(suite.T(), err.Error(), "token is expired")
}

func TestTokenSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
