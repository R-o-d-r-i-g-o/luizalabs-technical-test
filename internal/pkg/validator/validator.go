package validator

import (
	"luizalabs-technical-test/internal/pkg/formatter"
	"luizalabs-technical-test/pkg/constants/str"
)

// ValidateZipCode checks if the given string is a valid Brazilian ZipCode.
func ValidateZipCode(zipCode string) bool {
	zipCode = formatter.StripNonNumericCharacters(zipCode)
	return len(zipCode) == str.MinimumZipCodeLen
}
