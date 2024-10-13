package dependencies

import (
	"luizalabs-technical-test/internal/features/health"
	"luizalabs-technical-test/internal/features/swagger"
	"luizalabs-technical-test/internal/features/zipcode"
	"luizalabs-technical-test/internal/pkg/middleware"
	"luizalabs-technical-test/pkg/cache"
	"luizalabs-technical-test/pkg/http"
	"time"

	netHttp "net/http"

	"github.com/gin-gonic/gin"
)

const cleanupInterval = 1 * time.Minute

// Load sets up and returns a list of handler registration functions
func Load() []func(*gin.RouterGroup) {
	httpClient := http.NewClient(&netHttp.Client{})

	cacheManager := cache.NewManager(cleanupInterval)
	cacheMiddleware := middleware.NewCacheMiddleware(cacheManager)

	// zipcode feature
	zipCodeRep := zipcode.NewRepository(httpClient)
	zipCodeSrv := zipcode.NewService(zipCodeRep)
	zipCodeHandler := zipcode.NewHandler(zipCodeSrv, cacheMiddleware)

	// health feature
	healthHandler := health.NewHandler()

	// swagger feature
	swaggerHandler := swagger.NewHandler()

	return []func(*gin.RouterGroup){
		swaggerHandler.Register,
		healthHandler.Register,
		zipCodeHandler.Register,
	}
}
