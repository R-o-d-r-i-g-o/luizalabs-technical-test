package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// logger is the global zerolog instance, configured to log to stdout with timestamps.
var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

// Warn returns a "Warning" level log event.
func Warn(msg string) {
	logger.Info().Msg(msg)
}

// Debug returns a "Debug" level log event.
func Debug(msg string) {
	logger.Debug().Msg(msg)
}

// Error returns an "Error" level log event.
func Error(err error) {
	if err != nil {
		logger.Error().Err(err).Msg(err.Error())
	}
}
