package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// Now terminates the application immediately with an exit code of 1.
func Now() {
	os.Exit(1)
}

// Gracefully waits for termination signals and calls the cleanup function.
func Gracefully(app, cleanup func()) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go app()

	<-signalCh
	cleanup()
}
