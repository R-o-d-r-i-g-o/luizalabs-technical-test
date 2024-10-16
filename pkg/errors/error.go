package errors

import (
	"errors"
	"fmt"
	"luizalabs-technical-test/pkg/logger"
)

// ErrorImp is an interface implements error custom methods.
type ErrorImp interface {
	Error() string
	CodeStr() string
	WithErr(err error) *Error
	WithStrErr(format string, args ...interface{}) *Error
}

// Error represents a structured error with a code, message, and wrapped error.
type Error struct {
	Code    string
	Message string
	Err     error
}

// Error implements the error interface, returning the error message.
// This method ensures that only the custom error message is exposed to keep inside business rules out private and secure.
func (e *Error) Error() string {
	// If there is wrapped error, print it internally for future debug.
	if e.Err != nil {
		logger.Error(e.Err)
	}
	return e.Message
}

// CodeStr returns the error code as a string.
// This method is useful when you need to extract the error code.
func (e *Error) CodeStr() string {
	return e.Code
}

// WithErr returns a new Error instance with the provided underlying error.
func (e *Error) WithErr(err error) *Error {
	e.Err = err
	return e
}

// WithStrErr allows setting a formatted error message with the underlying error.
// It accepts a string format and arguments, formats the message, and wraps the resulting error.
func (e *Error) WithStrErr(format string, args ...interface{}) *Error {
	formattedMessage := fmt.Sprintf(format, args...)
	e.Err = errors.New(formattedMessage)
	return e
}
