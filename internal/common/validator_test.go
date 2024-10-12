package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateZipCode(t *testing.T) {
	testCases := []struct {
		zipCode  string
		expected bool
	}{
		{"12345-678", true},
		{"12345678", true},
		{"1234-5678", true},
		{"123456", false},
		{"12345-abc", false},
		{"12345-6789", false},
		{"   12345-678   ", true},
		{"   12345678   ", true},
		{"12345", false},
		{"1234567", false},
		{"", false},
	}

	// ASSERT
	for _, tt := range testCases {
		// ACT & ASSERT
		actual := ValidateZipCode(tt.zipCode)
		assert.Equal(t, tt.expected, actual)
	}
}
