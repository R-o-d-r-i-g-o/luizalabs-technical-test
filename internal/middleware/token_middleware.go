package middleware

import (
	"luizalabs-technical-test/internal/config"
	"luizalabs-technical-test/pkg/constants/str"
	"luizalabs-technical-test/pkg/middleware"
	"luizalabs-technical-test/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tokenMiddleware struct{}

func NewTokenMiddleware() middleware.Middleware {
	return &tokenMiddleware{}
}

func (t *tokenMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := token.ExtractBearerToken(c.Request)
		if tokenString == str.EMPTY_STRING {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims, err := token.ValidateToken(config.GeneralConfig.SECRET_AUTH_TOKEN_KEY, tokenString)
		if err != nil {
			// Customize the error message for unauthorized access
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set(token.CLAIMS_HEADER_NAME, claims)
		c.Next()
	}
}
