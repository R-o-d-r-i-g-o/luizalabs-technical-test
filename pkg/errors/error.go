package errors

import (
	"errors"
	"fmt"
)

// Error struct holds information about an error, allowing for detailed error handling and retrocompatibility
type Error struct {
	// Type represents the category or family of the error, useful for generalized error handling
	Type ErrorType `json:"type,omitempty"`

	// Context provides information on where the error occurred
	Context string `json:"context,omitempty"`

	// Name acts as the unique identifier for the error, enabling identification across different parts of the application
	Name string `json:"name,omitempty"`

	// Message is the main error message
	Message string `json:"message"`

	// Param is an optional parameter that gives extra information related to the error
	Param string `json:"param,omitempty"`

	// Errors is a list of nested errors, useful when an error contains multiple errors
	Errors []Error `json:"errors,omitempty"`

	// Err holds the Go error, allowing retrocompatibility with the standard error handling in Go
	err error

	// messageTemplate holds the template for the error message, which can be used with dynamic arguments
	messageTemplate string
}

// New creates and returns a new Error object with default values
func New() Error {
	return Error{}
}

// Error returns the error message, satisfying the built-in error interface
func (e Error) Error() string {
	return e.err.Error()
}

// Unwrap returns the underlying error for further inspection. Returns nil if no underlying error exists
func (e Error) Unwrap() error {
	return errors.Unwrap(e.err)
}

// Is compares the current error with the target error, checking for equality based on Name and Context
func (e Error) Is(target error) bool {
	if target == nil {
		return false
	}

	if err, ok := target.(Error); ok {
		return err.Name == e.Name && err.Context == e.Context
	}

	return errors.Is(e.err, target)
}

// WithType sets the error type and returns the updated Error object
func (e Error) WithType(typ ErrorType) Error {
	e.Type = typ
	return e
}

// WithContext sets the error context and returns the updated Error object
func (e Error) WithContext(ctx string) Error {
	e.Context = ctx
	return e
}

// WithName sets the error name and returns the updated Error object
func (e Error) WithName(name string) Error {
	e.Name = name
	return e
}

// WithParam sets an optional parameter related to the error and returns the updated Error object
func (e Error) WithParam(param string) Error {
	e.Param = param
	return e
}

// WithMessage sets the error message using a formatted string and optional arguments, then returns the updated Error object
func (e Error) WithMessage(message string, v ...any) Error {
	e.err = fmt.Errorf(message, v...)
	e.Message = e.err.Error()
	return e
}

// WithTemplate sets a template for the error message and immediately applies it to the error object
func (e Error) WithTemplate(template string) Error {
	e.messageTemplate = template
	return e.WithMessage(template)
}

// WithArgs constructs the final error message by applying dynamic arguments to the template message
// The args should be passed in the same order as specified in the template
func (e Error) WithArgs(args ...any) Error {
	return e.WithMessage(e.messageTemplate, args...)
}

// Add appends another error (or wraps a non-custom error) to the current Error object and returns the updated Error
func (e Error) Add(err error) Error {
	if err, ok := err.(Error); ok {
		e.Errors = append(e.Errors, err)
		return e
	}

	// If the provided error is not of type Error, wrap it with a default Unknown error
	e.Errors = append(e.Errors, Unknown.WithMessage("%w", err))
	return e
}
