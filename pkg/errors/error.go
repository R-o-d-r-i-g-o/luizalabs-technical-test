package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	// Type is the error family, that way we can generalize the error handling
	Type ErrorType `json:"type,omitempty"`

	// Context is where the error happened
	Context string `json:"context,omitempty"`

	// Name is the error identifier, with that, we can match the error in any place
	Name string `json:"name,omitempty"`

	Message string `json:"message"`
	Param   string `json:"param,omitempty"`

	// Errors is used in case of an Error tha have multiple errors inside
	Errors []Error `json:"errors,omitempty"`

	// Err is the go error to have retrocompatibility with the standard lib
	err error

	messageTemplate string
}

func New() Error {
	return Error{}
}

// Error return the error message and satisfy error interface
func (e Error) Error() string {
	return e.err.Error()
}

// Unwrap returns the original error. If not found, return nil
func (e Error) Unwrap() error {
	return errors.Unwrap(e.err)
}

// Is compare the errors.
func (e Error) Is(target error) bool {
	if target == nil {
		return false
	}

	if err, ok := target.(Error); ok {
		return err.Name == e.Name && err.Context == e.Context
	}

	return errors.Is(e.err, target)
}

func (e Error) WithType(typ ErrorType) Error {
	e.Type = typ
	return e
}

func (e Error) WithContext(ctx string) Error {
	e.Context = ctx
	return e
}

func (e Error) WithName(name string) Error {
	e.Name = name
	return e
}

func (e Error) WithParam(param string) Error {
	e.Param = param
	return e
}

func (e Error) WithMessage(message string, v ...any) Error {
	e.err = fmt.Errorf(message, v...)
	e.Message = e.err.Error()
	return e
}

// WithTemplate sets the message template to compose a default message with dynamic arguments
func (e Error) WithTemplate(template string) Error {
	e.messageTemplate = template
	return e.WithMessage(template)
}

// WithArgs create the final message composing the templateMessage with the given args
//
// the args must be passed in the same order that was setted in the template message
func (e Error) WithArgs(args ...any) Error {
	return e.WithMessage(e.messageTemplate, args...)
}

func (e Error) Add(err error) Error {
	if err, ok := err.(Error); ok {
		e.Errors = append(e.Errors, err)
		return e
	}

	e.Errors = append(e.Errors, Unknown.WithMessage("%w", err))
	return e
}