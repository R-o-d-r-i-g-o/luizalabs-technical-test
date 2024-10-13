package dependencies

import (
	"luizalabs-technical-test/internal/config"
	"luizalabs-technical-test/internal/features/health"
	"luizalabs-technical-test/internal/features/swagger"
	"luizalabs-technical-test/internal/features/zipcode"
	"luizalabs-technical-test/pkg/http"

	netHttp "net/http"

	"github.com/gin-gonic/gin"
)

// Load sets up and returns a list of handler registration functions
func Load() []func(*gin.RouterGroup) {
	config.LoadEnv()
	httpClient := http.NewClient(&netHttp.Client{})

	// zipcode feature
	zipCodeRep := zipcode.NewRepository(httpClient)
	zipCodeSrv := zipcode.NewService(zipCodeRep)
	zipCodeHandler := zipcode.NewHandler(zipCodeSrv)

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
