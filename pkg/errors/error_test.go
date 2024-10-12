package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	// ARRANGE
	var (
		mockError   = "Custom error message"
		customError = Error{
			Code:    "E001",
			Message: mockError,
		}
	)

	// ACT & ASSERT
	assert.Equal(t, mockError, customError.Error())
}

func TestError_WithErr(t *testing.T) {
	// ARRANGE
	var (
		wrappedErr    = errors.New("wrapped error")
		originalError = Error{
			Code:    "E002",
			Message: "Original error message",
		}
	)

	// ACT
	newError := originalError.WithErr(wrappedErr)

	// ASSERT
	assert.Equal(t, originalError.Code, newError.Code)
	assert.Equal(t, originalError.Message, newError.Message)
	assert.Equal(t, wrappedErr, newError.Err)
}
