package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorTypes(t *testing.T) {
	tests := []struct {
		typ   ErrorType
		value string
	}{
		{
			typ:   TypeValidation,
			value: "validation",
		},
		{
			typ:   TypeInvalidField,
			value: "invalid_field",
		},
		{
			typ:   TypeNotFound,
			value: "not_found",
		},
		{
			typ:   TypeBusinessRule,
			value: "business_rule",
		},
		{
			typ:   TypeUnknown,
			value: "unknown",
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			assert.Equal(t, string(tt.typ), tt.value)
		})
	}
}

func TestTypeError_Fail(t *testing.T) {
	t.Run("should return false", func(t *testing.T) {
		mockedString := "I am an error"
		mockedError := errors.New(mockedString)

		givenErrorType := ErrorType(mockedString)
		receivedBoolean := givenErrorType.Match(mockedError)

		assert.False(t, receivedBoolean)
	})
}