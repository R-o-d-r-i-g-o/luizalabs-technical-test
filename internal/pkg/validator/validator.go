package validator

import "luizalabs-technical-test/internal/pkg/formatter"

// ValidateZipCode checks if the given string is a valid Brazilian ZipCode.
func ValidateZipCode(zipCode string) bool {
	zipCode = formatter.StripNonNumericCharacters(zipCode)
	return len(zipCode) == MinimumZipCodeLen
}
