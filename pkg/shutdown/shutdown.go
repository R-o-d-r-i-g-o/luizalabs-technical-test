package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// Now terminates the application immediately, allowing for potential recovery actions.
func Now() {
	panic("program has being terminated.")
}

// Gracefully waits for termination signals and calls the cleanup function.
func Gracefully(app, cleanup func()) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go app()

	<-signalCh
	cleanup()
}
