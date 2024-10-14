package logger

import (
	"bytes"
	"errors"
	"luizalabs-technical-test/pkg/constants/str"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// TestLogger captures logs in a buffer for testing.
func TestLogger(t *testing.T) {
	// Note: Replace the default logger with one that writes to the buffer.
	var buf bytes.Buffer
	logger = zerolog.New(&buf).With().Timestamp().Logger()

	// ARRANGE
	tests := []struct {
		name    string
		logFunc func(string)
		message string
	}{
		{"Warn", func(msg string) { Warn(msg) }, "This is a warning"},
		{"Debug", func(msg string) { Debug(msg) }, "This is a debug message"},
		{"Error", func(msg string) { Error(nil) }, str.EmptyString},
		{"Error", func(msg string) { Error(errors.New(msg)) }, "This is an error"},
	}

	// ACT & ASSERT
	for _, test := range tests {
		buf.Reset()
		test.logFunc(test.message)
		assert.Contains(t, buf.String(), test.message)
	}
}
