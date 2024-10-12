package common

import "regexp"

// AdjustLastNonZeroDigit replaces the first non-zero digit in the ZipCode string
// (traversing from the end) with '0'. Returns the original ZipCode if non-zero digits are found.
func AdjustLastNonZeroDigit(zipCode string) string {
	zipCode = StripNonNumericCharacters(zipCode)
	runes := []rune(zipCode)

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] != '0' {
			runes[i] = '0'
			break
		}
	}
	return string(runes)
}

// StripNonNumericCharacters removes all non-numeric characters from the input string.
func StripNonNumericCharacters(input string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(input, "")
}
