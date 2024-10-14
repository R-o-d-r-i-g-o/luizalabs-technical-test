package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"luizalabs-technical-test/internal/config"
	"luizalabs-technical-test/pkg/cache"
	"luizalabs-technical-test/pkg/env"
	"luizalabs-technical-test/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CacheMiddlewareSuite struct {
	suite.Suite
	cacheManager cache.Manager
}

func (s *CacheMiddlewareSuite) SetupSuite() {
	// Initialize the global cache manager for testing
	os.Setenv("SECRET_AUTH_TOKEN_KEY", "test-secret")

	env.LoadStructWithEnvVars("env", &config.ServerConfig, &config.GeneralConfig, &config.PostgresConfig)
	s.cacheManager = cache.NewManager(time.Minute)
}

func (s *CacheMiddlewareSuite) TestCacheMiddleware() {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	secretKey := "test-secret"
	userToken, err := token.CreateToken(secretKey, token.CustomClaims{CustomKeys: map[string]any{"Email": "test"}})
	assert.NoError(s.T(), err)

	// Create an instance of the cache middleware
	cacheMiddleware := NewCacheMiddleware(s.cacheManager)
	router.Use(cacheMiddleware.Middleware())
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// Test 1: Cache Miss
	req1, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req1.Header.Set("Authorization", "Bearer "+userToken)
	recorder1 := httptest.NewRecorder()
	router.ServeHTTP(recorder1, req1)

	assert.Equal(s.T(), http.StatusOK, recorder1.Code)
	var response1 map[string]interface{}
	err = json.Unmarshal(recorder1.Body.Bytes(), &response1)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "Hello, World!", response1["message"])

	// Test 2: Cache Hit
	recorder2 := httptest.NewRecorder()
	router.ServeHTTP(recorder2, req1) // Send the same request again

	assert.Equal(s.T(), http.StatusFound, recorder2.Code)
	var response2 map[string]interface{}
	err = json.Unmarshal(recorder2.Body.Bytes(), &response2)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "Hello, World!", response2["message"])
}

func (s *CacheMiddlewareSuite) TestCacheMiddlewareInvalidToken() {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	cacheMiddleware := NewCacheMiddleware(s.cacheManager)
	router.Use(cacheMiddleware.Middleware())
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// Mocking an invalid token extraction

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(s.T(), http.StatusOK, recorder.Code)
	var response map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "Hello, World!", response["message"])
}

func TestCacheMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(CacheMiddlewareSuite))
}
