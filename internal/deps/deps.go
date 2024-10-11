package deps

import (
	"luizalabs-technical-test/internal/features/cep"
	"luizalabs-technical-test/internal/features/health"

	"github.com/gin-gonic/gin"
)

func LoadDependencies() []func(*gin.Engine) {
	// cep feature
	cepRep := cep.NewRepository()
	cepSrv := cep.NewService(cepRep)
	cepHandler := cep.NewHandler(cepSrv)

	// health feature
	healthHandler := health.NewHandler()

	return []func(*gin.Engine){
		healthHandler.Register,
		cepHandler.Register,
	}
}
