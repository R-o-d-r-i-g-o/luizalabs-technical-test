package main

import (
	"fmt"

	"luizalabs-technical-test/internal/config"
	"luizalabs-technical-test/internal/dependencies"
	"luizalabs-technical-test/internal/pkg/cors"
	"luizalabs-technical-test/pkg/server"
	"luizalabs-technical-test/pkg/shutdown"
)

func main() {
	cleanup := func() {
		fmt.Println("Cleaning up resources...")

		// Add pull conn. close logic here

		fmt.Println("Cleanup done.")
	}

	runnapp := func() {
		srv := server.NewGinServer()

		srv.SetupCustom(cors.RouteSettings)
		srv.SetupHandlers("v1", dependencies.Load()...)
		srv.SetupMiddleware(cors.Middleware())

		err := srv.Run(":" + config.ServerConfig.Port)
		if err != nil {
			fmt.Printf("Error running server: %v\n", err)
			shutdown.Now()
		}
	}

	shutdown.Gracefully(runnapp, cleanup)
}
