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

type CustomClaims struct {
	jwt.StandardClaims
	CustomKeys map[string]any `json:"custom_claims,omitempty"`
}

const CLAIMS_HEADER_NAME = "claims"

func CreateToken(secretKey string, claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return str.EMPTY_STRING, err
	}

	return tokenString, nil
}

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

func ExtractTokenClaimsFromContext(secretKey string, ctx context.Context) (CustomClaims, error) {
	ginContext, ok := ctx.(*gin.Context)
	if !ok {
		return CustomClaims{}, errors.New("failed to parse context to gin context")
	}

	token := ExtractBearerToken(ginContext.Request)
	if token == str.EMPTY_STRING {
		return CustomClaims{}, errors.New("token not found")
	}

	tokenClaims, err := ValidateToken(secretKey, token)
	if err != nil {
		return CustomClaims{}, err
	}

	return *tokenClaims, nil
}

func ExtractBearerToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == str.EMPTY_STRING {
		return str.EMPTY_STRING
	}

	return strings.Split(authHeader, str.EMPTY_SPACE)[1]
}
