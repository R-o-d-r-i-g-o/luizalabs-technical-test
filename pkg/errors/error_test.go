package errors

import (
	"testing"

	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ErrorSuite defines a test suite for the error structure.
type ErrorSuite struct {
	suite.Suite
}

// TestCodeStr tests the implementation of the CodeStr method.
func (suite *ErrorSuite) TestCodeStr() {
	tests := []struct {
		name         string
		customErr    *Error
		expectedCode string
	}{
		{
			name:         "With valid code",
			customErr:    &Error{Code: "400", Message: "custom error"},
			expectedCode: "400",
		},
		{
			name:         "With empty code",
			customErr:    &Error{Code: "", Message: "custom error"},
			expectedCode: "",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			assert.Equal(suite.T(), tt.expectedCode, tt.customErr.CodeStr())
		})
	}
}

// TestError tests the implementation of the Error method.
func (suite *ErrorSuite) TestError() {
	tests := []struct {
		name        string
		customErr   *Error
		expectedMsg string
	}{
		{
			name:        "With wrapped error",
			customErr:   &Error{Code: "400", Message: "custom error", Err: errors.New("internal error")},
			expectedMsg: "custom error",
		},
		{
			name:        "Without wrapped error",
			customErr:   &Error{Code: "400", Message: "custom error"},
			expectedMsg: "custom error",
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			assert.Equal(suite.T(), tt.expectedMsg, tt.customErr.Error())
		})
	}
}

// TestWithErr tests the WithErr method.
func (suite *ErrorSuite) TestWithErr() {
	customErr := &Error{
		Code:    "400",
		Message: "custom error",
	}

	subErr := errors.New("another error")
	newErr := customErr.WithErr(subErr)

	assert.Equal(suite.T(), subErr, newErr.Err)
}

// TestWithStrErr tests the WithStrErr method.
func (suite *ErrorSuite) TestWithStrErr() {
	customErr := &Error{
		Code:    "400",
		Message: "custom error",
	}

	formattedErr := customErr.WithStrErr("formatted error: %s", "details")

	assert.Equal(suite.T(), "formatted error: details", formattedErr.Err.Error())
	assert.Equal(suite.T(), "custom error", formattedErr.Error())
}

// TestMain runs the test suite.
func TestErrorSuite(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}
