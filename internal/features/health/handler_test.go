package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	handler := NewHandler()
	handler.Register(router.Group("/v1"))

	// Test /ping endpoint
	t.Run("GET /ping", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/v1/health/ping", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		// Check the response body
		var response swagHealthResponse
		err := json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "pong", response.Data.Message)
	})

	// Test /metrics endpoint
	t.Run("GET /metrics", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/v1/health/metrics", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "# HELP") // Check for Prometheus metrics output
	})
}
