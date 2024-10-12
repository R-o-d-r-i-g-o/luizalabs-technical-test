package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonErrors(t *testing.T) {
	tests := []struct {
		typ ErrorType
		err Error
	}{
		{
			typ: TypeBusinessRule,
			err: BusinessRule,
		},
		{
			typ: TypeInvalidField,
			err: InvalidField,
		},
		{
			typ: TypeNotFound,
			err: NotFound,
		},
		{
			typ: TypeValidation,
			err: Validation,
		},
		{
			typ: TypeUnknown,
			err: Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.typ), func(t *testing.T) {
			assert.True(t, tt.typ.Match(tt.err))
		})
	}

}