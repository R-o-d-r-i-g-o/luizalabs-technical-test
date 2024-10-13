package token

import (
	"context"
	"errors"
	"luizalabs-technical-test/pkg/constants/str"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CustomClaims represents a custom set of claims embedded in the JWT token.
// It extends the standard JWT claims (e.g., expiration, issuer) with a map of additional custom keys.
type CustomClaims struct {
	jwt.StandardClaims
	CustomKeys map[string]any `json:"custom_claims,omitempty"`
}

// ClaimsHeaderName is the key used to store token claims in the Gin context.
const ClaimsHeaderName = "claims"

// CreateToken generates a JWT token using a secret key and custom claims.
func CreateToken(secretKey string, claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return str.EmptyString, err
	}

	return tokenString, nil
}

// ValidateToken verifies the provided token string using the secret key and returns the custom claims if valid.
func ValidateToken(secretKey, tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ExtractTokenClaimsFromContext extracts and validates the JWT token from the Gin context and returns the custom claims.
func ExtractTokenClaimsFromContext(ctx context.Context, secretKey string) (CustomClaims, error) {
	ginContext, ok := ctx.(*gin.Context)
	if !ok {
		return CustomClaims{}, errors.New("failed to parse context to gin context")
	}

	token := ExtractBearerToken(ginContext.Request)
	if token == str.EmptyString {
		return CustomClaims{}, errors.New("token not found")
	}

	tokenClaims, err := ValidateToken(secretKey, token)
	if err != nil {
		return CustomClaims{}, err
	}

	return *tokenClaims, nil
}

// ExtractBearerToken extracts the Bearer token from the Authorization header of the HTTP request.
func ExtractBearerToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == str.EmptyString {
		return str.EmptyString
	}

	return strings.Split(authHeader, str.EmptySpace)[1]
}
