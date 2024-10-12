package common

// ValidateZipCode checks if the given string is a valid Brazilian ZipCode.
func ValidateZipCode(zipCode string) bool {
	zipCode = StripNonNumericCharacters(zipCode)
	return len(zipCode) == MinimumZipCodeLen
}
