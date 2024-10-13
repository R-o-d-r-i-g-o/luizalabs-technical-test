package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"luizalabs-technical-test/internal/config"
	"luizalabs-technical-test/pkg/middleware"
	"luizalabs-technical-test/pkg/token"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TokenMiddlewareTestSuite struct {
	suite.Suite
	router *gin.Engine
	mw     middleware.Middleware
}

func (suite *TokenMiddlewareTestSuite) SetupSuite() {
	// Set up Gin router and middleware
	suite.router = gin.New()
	suite.mw = NewTokenMiddleware()
	suite.router.Use(suite.mw.Middleware())
	suite.router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Set up config for the tests
	config.GeneralConfigEnv.SecretAuthTokenKey = "test_secret_key"
}

func (suite *TokenMiddlewareTestSuite) TestTokenMiddleware() {
	tests := []struct {
		name         string
		authHeader   string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "No token provided",
			authHeader:   "",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"error":"Unauthorized"}`,
		},
		{
			name:         "Invalid token provided",
			authHeader:   "Bearer invalidtoken",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"error":"invalid token"}`, // Adjust based on your error handling
		},
		{
			name:         "Valid token provided",
			authHeader:   "Bearer " + suite.createValidToken(),
			expectedCode: http.StatusOK,
			expectedBody: `{"message":"success"}`,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			// Create a request with the appropriate Authorization header
			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			// Create a response recorder
			w := httptest.NewRecorder()
			suite.router.ServeHTTP(w, req)

			// Assert the response
			assert.Equal(suite.T(), tt.expectedCode, w.Code)
			assert.JSONEq(suite.T(), tt.expectedBody, w.Body.String())
		})
	}
}

// Helper function to create a valid token
func (suite *TokenMiddlewareTestSuite) createValidToken() string {
	claims := token.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   "1234567890",
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		CustomKeys: map[string]any{"foo": "bar"},
	}
	tokenString, _ := token.CreateToken(config.GeneralConfigEnv.SecretAuthTokenKey, claims)
	return tokenString
}

func TestTokenMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(TokenMiddlewareTestSuite))
}
