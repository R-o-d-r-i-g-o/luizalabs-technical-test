package middleware

import (
	"luizalabs-technical-test/internal/config"
	"luizalabs-technical-test/pkg/constants/str"
	"luizalabs-technical-test/pkg/middleware"
	"luizalabs-technical-test/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TokenMiddleware is an interface that extends the base middleware.Middleware interface.
// It is used to define middleware logic for handling authentication tokens, enabling
// token validation and authorization checks in the request flow.
type TokenMiddleware interface {
	middleware.Middleware
}

type tokenMiddleware struct{}

// NewTokenMiddleware creates a new instance of tokenMiddleware, which validates tokens for authentication
func NewTokenMiddleware() TokenMiddleware {
	return &tokenMiddleware{}
}

// Middleware validates the Bearer token in incoming requests. If valid, the token claims are added to the context.
// If the token is missing or invalid, it aborts the request with an unauthorized status.
func (t *tokenMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := token.ExtractBearerToken(c.Request)
		if tokenString == str.EmptyString {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims, err := token.ValidateToken(config.GeneralConfig.SecretAuthTokenKey, tokenString)
		if err != nil {
			// Customize the error message for unauthorized access
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Set token claims in the context for further use in the request lifecycle
		c.Set(token.ClaimsHeaderName, claims)
		c.Next()
	}
}
