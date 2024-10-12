package common

// ValidateZipCode checks if the given string is a valid Brazilian ZipCode.
func ValidateZipCode(zipCode string) bool {
	const _MINIMUN_ZIPCODE_LEN = 8
	zipCode = StripNonNumericCharacters(zipCode)

	return len(zipCode) == _MINIMUN_ZIPCODE_LEN
}
