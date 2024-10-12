package main

import (
	"fmt"

	"luizalabs-technical-test/internal/deps"
	"luizalabs-technical-test/pkg/server"
	"luizalabs-technical-test/pkg/shutdown"
)

func init() {}

func main() {
	cleanup := func() {
		fmt.Println("Cleaning up resources...")

		// Add pull conn. close logic here

		fmt.Println("Cleanup done.")
	}

	runnapp := func() {
		srv := server.NewGinServer()
		srv.SetupHandlers(deps.LoadDependencies()...)

		err := srv.Run(":8080")
		if err != nil {
			fmt.Printf("Error running server: %v\n", err)
			shutdown.Now()
		}
	}

	shutdown.Gracefully(runnapp, cleanup)
}
