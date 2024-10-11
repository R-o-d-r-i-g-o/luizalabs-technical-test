package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// GracefulShutdown waits for termination signals and calls the cleanup function.
func GracefulShutdown(app, cleanup func()) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go app()

	<-signalCh
	cleanup()
}
