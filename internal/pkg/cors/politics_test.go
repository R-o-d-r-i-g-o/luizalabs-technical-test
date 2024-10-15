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

	tests := []struct {
		name         string
		method       string
		expectedCode int
	}{
		{
			name:         "Preflight Request",
			method:       http.MethodOptions,
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "Successful Request",
			method:       http.MethodGet,
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, "/test", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestRouteSettings(t *testing.T) {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "sample end-point"})
	})
	RouteSettings(router)

	tests := []struct {
		name         string
		method       string
		route        string
		expectedCode int
	}{
		{
			name:         "Non-Existing Route",
			method:       http.MethodGet,
			route:        "/non-existing",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Method Not Allowed",
			method:       http.MethodPost,
			route:        "/",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "Successful Request",
			method:       http.MethodGet,
			route:        "/",
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.route, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
