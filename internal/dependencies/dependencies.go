package dependencies

import (
	"luizalabs-technical-test/internal/features/cep"
	"luizalabs-technical-test/internal/features/health"
	"luizalabs-technical-test/internal/features/swagger"
	"luizalabs-technical-test/pkg/http"

	netHttp "net/http"

	"github.com/gin-gonic/gin"
)

// Load sets up and returns a list of handler registration functions
func Load() []func(*gin.RouterGroup) {
	httpClient := http.NewClient(&netHttp.Client{})

	// cep feature
	cepRep := cep.NewRepository(httpClient)
	cepSrv := cep.NewService(cepRep)
	cepHandler := cep.NewHandler(cepSrv)

	// health feature
	healthHandler := health.NewHandler()

	// swagger feature
	swaggerHandler := swagger.NewHandler()

	return []func(*gin.RouterGroup){
		swaggerHandler.Register,
		healthHandler.Register,
		cepHandler.Register,
	}
}
