package deps

import (
	"luizalabs-technical-test/internal/features/cep"
	"luizalabs-technical-test/internal/features/health"
	"luizalabs-technical-test/internal/features/swagger"

	"github.com/gin-gonic/gin"
)

// LoadDependencies sets up and returns a list of handler registration functions
func LoadDependencies() []func(*gin.Engine) {
	// cep feature
	cepRep := cep.NewRepository()
	cepSrv := cep.NewService(cepRep)
	cepHandler := cep.NewHandler(cepSrv)

	// health feature
	healthHandler := health.NewHandler()

	// swagger feature
	swaggerHandler := swagger.NewHandler()

	return []func(*gin.Engine){
		swaggerHandler.Register,
		healthHandler.Register,
		cepHandler.Register,
	}
}
