package main

import (
	"luizalabs-technical-test/internal/config"
	"luizalabs-technical-test/internal/dependencies"
	"luizalabs-technical-test/internal/pkg/cors"
	"luizalabs-technical-test/pkg/logger"
	"luizalabs-technical-test/pkg/postgres"
	"luizalabs-technical-test/pkg/server"
	"luizalabs-technical-test/pkg/shutdown"
)

func main() {
	cleanup := func() {
		logger.Warn("service stop running...")
		postgres.Close()
		logger.Warn("server stoped correctly.")
	}

	runnapp := func() {
		srv := server.NewGinServer()

		srv.SetupCustom(cors.RouteSettings)
		srv.SetupHandlers("v1", dependencies.Load()...)
		srv.SetupMiddleware(cors.Middleware())
		logger.Warn("starting server on port: " + config.ServerConfig.Port)

		err := srv.Run(":" + config.ServerConfig.Port)
		if err != nil {
			logger.Error(err)
			shutdown.Now()
		}
	}

	shutdown.Gracefully(runnapp, cleanup)
}
