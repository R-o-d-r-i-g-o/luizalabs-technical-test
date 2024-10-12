package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjustLastNonZeroDigit(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"12345", "12340"},   // Normal case
		{"67890", "67800"},   // Replace last non-zero digit
		{"00000", "00000"},   // All zeros
		{"10001", "10000"},   // Single non-zero at the end
		{"54321", "54320"},   // Multiple non-zeros
		{"01234", "01230"},   // Leading zero case
		{"00001", "00000"},   // Single non-zero at the end
		{"987650", "987600"}, // No change with last digit zero
		{"12-345", "12340"},  // Case with special character
		{"12 345", "12340"},  // Case with space
		{"12A45", "1240"},    // Case with non-numeric character
	}

	// ARRANGE
	for _, tt := range testCases {
		// ACT & ASSERT
		actual := AdjustLastNonZeroDigit(tt.input)
		assert.Equal(t, tt.expected, actual)

	}
}

func TestStripNonNumericCharacters(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"12345", "12345"},               // All numeric
		{"12a34b5", "12345"},             // Alphanumeric
		{"!@#$%^&*()_+", ""},             // Special characters only
		{"CEP: 12345-6789", "123456789"}, // Mixed with text and symbols
		{"abc123xyz", "123"},             // Letters with numbers
		{"0x123", "0123"},                // Hexadecimal prefix
		{"", ""},                         // Empty string
		{"   123   ", "123"},             // Spaces around numbers
	}

	// ARRANGE
	for _, tt := range tests {
		// ACT & ASSERT
		actual := StripNonNumericCharacters(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}
