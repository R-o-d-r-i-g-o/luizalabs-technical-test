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
	req, _ := http.NewRequest(http.MethodOptions, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code, "Expected status code 204 for OPTIONS request")
}

func TestRouteSettings(t *testing.T) {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "sample end-point"})
	})

	RouteSettings(router)

	// Test non-existing route
	req, _ := http.NewRequest(http.MethodGet, "/non-existing", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	expectedError := `{"code":"ERR_NO_ROUTE","error":"A rota solicitada não foi encontrada. Por favor, verifique a URL e tente novamente."}`
	assert.JSONEq(t, expectedError, w.Body.String(), "Expected error response for no route found")

	// Test method not allowed
	req, _ = http.NewRequest(http.MethodPost, "/", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	expectedMethodError := `{"code":"ERR_NO_METHOD","error":"O método solicitado não é permitido para esta rota. Por favor, verifique os métodos permitidos e tente novamente."}`
	assert.JSONEq(t, expectedMethodError, w.Body.String(), "Expected error response for method not allowed")
}
