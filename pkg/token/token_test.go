package token_test

import (
	"context"
	"luizalabs-technical-test/pkg/constants/str"
	"luizalabs-technical-test/pkg/token"
	"net/http"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

// Test for ExtractTokenClaimsFromContext
func (suite *TokenTestSuite) TestExtractTokenClaimsFromContext() {
	// Set up the gin context and token
	secretKey := "secret_key"
	invalidToken := "invalidToken"

	tests := []struct {
		name           string
		mockContext    context.Context
		expectedClaims token.CustomClaims
		expectError    bool
	}{
		{
			name:           "Failed to parse context",
			mockContext:    context.Background(),
			expectedClaims: token.CustomClaims{},
			expectError:    true,
		},
		{
			name:           "Token not found",
			mockContext:    createGinContextWithToken(""),
			expectedClaims: token.CustomClaims{},
			expectError:    true,
		},
		{
			name:           "Error during token validation",
			mockContext:    createGinContextWithToken(invalidToken),
			expectedClaims: token.CustomClaims{},
			expectError:    true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			claims, err := token.ExtractTokenClaimsFromContext(tt.mockContext, secretKey)

			if tt.expectError {
				assert.Error(suite.T(), err)
			} else {
				assert.NoError(suite.T(), err)
				assert.Equal(suite.T(), tt.expectedClaims, claims)
			}
		})
	}
}

func (suite *TokenTestSuite) TestExtractBearerToken() {
	tests := []struct {
		name          string
		authHeader    string
		expectedToken string
	}{
		{
			name:          "Valid Bearer Token",
			authHeader:    "Bearer test_token",
			expectedToken: "test_token",
		},
		{
			name:          "Bearer Token Without Prefix",
			authHeader:    "test_token",
			expectedToken: "test_token",
		},
		{
			name:          "Empty Authorization Header",
			authHeader:    str.EmptyString,
			expectedToken: str.EmptyString,
		},
		{
			name:          "Invalid Authorization Format",
			authHeader:    "InvalidHeader test_token",
			expectedToken: "test_token",
		},
		{
			name:          "Only Bearer Without Token",
			authHeader:    "Bearer ",
			expectedToken: str.EmptyString,
		},
		{
			name:          "Spaces Only",
			authHeader:    str.EmptySpace,
			expectedToken: str.EmptyString,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Authorization", tt.authHeader)

			token := token.ExtractBearerToken(req)
			assert.Equal(suite.T(), tt.expectedToken, token)
		})
	}
}

// Helper function to create a gin context with a mock request containing a bearer token
func createGinContextWithToken(token string) context.Context {
	ginContext := &gin.Context{
		Request: &http.Request{
			Header: http.Header{
				"Authorization": []string{"Bearer " + token},
			},
		},
	}

	return ginContext
}

func TestTokenSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
