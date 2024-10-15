package swagger

import (
	"luizalabs-technical-test/docs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSwaggerHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	handler := NewHandler()
	handler.Register(router.Group("/v1"))

	// Test Swagger Documentation Endpoint
	t.Run("GET /docs", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/v1/docs", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusMovedPermanently, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Moved Permanently")
	})

	// Test Swagger Documentation with Wildcard
	t.Run("GET /docs/*any", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/v1/docs/index.html", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Not Found")
	})

	// Test swagger custom settings (dinamically setted)
	t.Run("Check for custom settings", func(t *testing.T) {
		assert.Equal(t, "localhost:8080", docs.SwaggerInfo.Host)
		assert.Equal(t, "luizalabs-technical-test", docs.SwaggerInfo.Title)
		assert.Equal(t, "1.0", docs.SwaggerInfo.Version)
		assert.Equal(t, "/", docs.SwaggerInfo.BasePath)
		assert.Equal(t, "Clear and concise documentation detailing each API route implementation.", docs.SwaggerInfo.Description)
		assert.ElementsMatch(t, []string{"http", "https"}, docs.SwaggerInfo.Schemes)
	})
}
