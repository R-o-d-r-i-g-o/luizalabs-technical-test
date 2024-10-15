package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"luizalabs-technical-test/internal/config"
	"luizalabs-technical-test/pkg/cache"
	"luizalabs-technical-test/pkg/constants/str"
	"luizalabs-technical-test/pkg/middleware"
	"luizalabs-technical-test/pkg/token"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	defaultCacheTimeout time.Duration = time.Minute * 30
	cacheHeaderKey      string        = "X-Cache-Control"
	noCache             string        = "no-cache"
)

type cacheMiddleware struct {
	cacheManager cache.Manager
}

type cacheWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// NewCacheMiddleware creates a new instance of the cache middleware
func NewCacheMiddleware(cacheManager cache.Manager) middleware.Middleware {
	return &cacheMiddleware{cacheManager: cacheManager}
}

// Middleware is the function that provides the cache middleware logic to be used in the Gin router.
// It checks for a cache hit, and if found, returns the cached value. Otherwise, it processes the request and caches the response.
func (c *cacheMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !c.shouldHandleRequest(ctx) {
			ctx.Next()
			return
		}

		cacheKey, err := c.createCacheKeyFromRequest(ctx)
		if err != nil {
			ctx.Next()
			return
		}

		if cache, exists := c.cacheManager.Get(cacheKey); exists {
			ctx.AbortWithStatusJSON(http.StatusOK, c.parseCachedValue(cache))
			return
		}

		cacheWriter := newCacheWriter(ctx.Writer)
		ctx.Writer = cacheWriter

		ctx.Next()
		c.cacheResponse(cacheKey, cacheWriter.body.Bytes())
	}
}

// NewCacheWriter creates a new cache writer that intercepts the response body for caching
func newCacheWriter(w gin.ResponseWriter) *cacheWriter {
	return &cacheWriter{ResponseWriter: w, body: bytes.NewBuffer(nil)}
}

// Write intercepts the response body and writes both to the buffer and the original response writer
func (w *cacheWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

// ShouldHandleRequest checks if the request method is GET and if caching is allowed based on request headers
func (c *cacheMiddleware) shouldHandleRequest(ctx *gin.Context) bool {
	if ctx.Request.Method != http.MethodGet {
		return false
	}

	value := ctx.GetHeader(cacheHeaderKey)
	return value == str.EmptyString || strings.ToLower(value) == noCache
}

// ParseCachedValue attempts to unmarshal cached data into either a map or an array and returns the parsed value
func (c *cacheMiddleware) parseCachedValue(value interface{}) interface{} {
	var dataMap map[string]interface{}
	if err := json.Unmarshal(value.([]byte), &dataMap); err == nil {
		return dataMap
	}

	var dataArray []map[string]interface{}
	if err := json.Unmarshal(value.([]byte), &dataArray); err == nil {
		return dataArray
	}

	return nil
}

// CacheResponse stores the response body in the cache with a predefined timeout
func (c *cacheMiddleware) cacheResponse(key string, body []byte) {
	c.cacheManager.Set(key, body, defaultCacheTimeout)
}

// CreateCacheKeyFromRequest generates a cache key using the user's hash and the request URL
func (c *cacheMiddleware) createCacheKeyFromRequest(ctx *gin.Context) (string, error) {
	userHash, err := c.getUserHashFromTokenClaims(ctx)
	if err != nil {
		return str.EmptyString, err
	}

	return fmt.Sprintf("%s%s", userHash, ctx.Request.URL), nil
}

// GetUserHashFromTokenClaims extracts the user hash from token claims in the request context
func (c *cacheMiddleware) getUserHashFromTokenClaims(ctx *gin.Context) (string, error) {
	claims, err := token.ExtractTokenClaimsFromContext(ctx, config.GeneralConfig.SecretAuthTokenKey)
	if err != nil {
		return str.EmptyString, err
	}

	hash, ok := claims.CustomKeys["Email"].(string)
	if !ok {
		return hash, errors.New("no hash found")
	}

	return hash, nil
}
