package cors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	router := gin.New()
	router.Use(Middleware())
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Test preflight request
	t.Run("Preflight Request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodOptions, "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code, "Expected status code 204 for OPTIONS request")
	})
}

func TestRouteSettings(t *testing.T) {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "sample end-point"})
	})
	RouteSettings(router)

	t.Run("Non-Existing Route", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/non-existing", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code, "Expected status code 404 for non-existing route")
	})

	t.Run("Method Not Allowed", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusMethodNotAllowed, w.Code, "Expected status code 405 for method not allowed")
	})
}
