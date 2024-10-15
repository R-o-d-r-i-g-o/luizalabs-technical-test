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
	"luizalabs-technical-test/pkg/crypt"
	"luizalabs-technical-test/pkg/http"
	"luizalabs-technical-test/pkg/logger"
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
	cryptHasher := crypt.NewPasswordHasher()
	logger.Debug("Instanciate internal dependencies...")

	cacheManager := cache.NewManager(cleanupInterval)
	cacheMiddleware := middleware.NewCacheMiddleware(cacheManager)
	tokenMiddleware := middleware.NewTokenMiddleware()
	logger.Debug("Instanciate middleware dependencies...")

	// auth feature
	authRep := auth.NewRepository(db)
	authSrv := auth.NewService(authRep, cryptHasher)
	authHandler := auth.NewHandler(authSrv)
	logger.Debug("Instanciate auth use-case dependencies...")

	// zipcode feature
	zipCodeRep := zipcode.NewRepository(httpClient)
	zipCodeSrv := zipcode.NewService(zipCodeRep)
	zipCodeHandler := zipcode.NewHandler(zipCodeSrv, cacheMiddleware, tokenMiddleware)
	logger.Debug("Instanciate zipcode use-case dependencies...")

	// health feature
	healthHandler := health.NewHandler()
	logger.Debug("Instanciate health use-case dependencies...")

	// swagger feature
	swaggerHandler := swagger.NewHandler()
	logger.Debug("Instanciate swagger use-case dependencies...")

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
		logger.Error(err)
		shutdown.Now()
	}

	postgres.Migrate(entity.User{})
	return db
}
