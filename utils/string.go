package utils

import "unicode"

// counts the number of alphanumeric characters in a string
func CountAlphaNumeric(s string) int {
	count := 0

	// iterate over each character in the string
	for _, c := range s {
		// check if the character is a letter or a number
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			// increment the count if the character is alphanumeric
			count++
		}
	}

	return count
}
