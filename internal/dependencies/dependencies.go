package dependencies

import (
	"luizalabs-technical-test/internal/config"
	"luizalabs-technical-test/internal/features/auth"
	"luizalabs-technical-test/internal/features/health"
	"luizalabs-technical-test/internal/features/swagger"
	"luizalabs-technical-test/internal/features/zipcode"
	"luizalabs-technical-test/internal/pkg/entity"
	"luizalabs-technical-test/internal/pkg/middleware"
	"luizalabs-technical-test/pkg/cache"
	"luizalabs-technical-test/pkg/http"
	"luizalabs-technical-test/pkg/postgres"
	"luizalabs-technical-test/pkg/shutdown"
	"time"

	netHttp "net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const cleanupInterval = 1 * time.Minute

// Load sets up and returns a list of handler registration functions
func Load() []func(*gin.RouterGroup) {
	db := loadPostgresDepencies()
	httpClient := http.NewClient(&netHttp.Client{})

	cacheManager := cache.NewManager(cleanupInterval)
	cacheMiddleware := middleware.NewCacheMiddleware(cacheManager)

	// auth feature
	authRep := auth.NewRepository(db)
	authSrv := auth.NewService(authRep)
	authHandler := auth.NewHandler(authSrv)

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
		authHandler.Register,
	}
}

func loadPostgresDepencies() *gorm.DB {
	postgres.SetConnectionString(config.PostgresConfig.ToPostgresDSN())
	db, err := postgres.GetInstance()
	if err != nil {
		shutdown.Now()
	}

	postgres.Migrate(entity.User{})
	return db
}
