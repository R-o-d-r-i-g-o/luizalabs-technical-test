package common

import (
	"regexp"
)

// ValidateZipCode checks if the given string is a valid Brazilian ZipCode.
func ValidateZipCode(zipCode string) bool {
	zipCode = StripNonNumericCharacters(zipCode)

	if len(zipCode) != 8 {
		return false
	}

	re := regexp.MustCompile(`^\d{5}(-?\d{3})?$`)
	return re.MatchString(zipCode)
}
