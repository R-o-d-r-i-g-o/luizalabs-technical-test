package dependencies

import (
	"luizalabs-technical-test/internal/features/cep"
	"luizalabs-technical-test/internal/features/health"
	"luizalabs-technical-test/internal/features/swagger"

	"github.com/gin-gonic/gin"
)

// Load sets up and returns a list of handler registration functions
func Load() []func(*gin.RouterGroup) {
	// cep feature
	cepRep := cep.NewRepository()
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
