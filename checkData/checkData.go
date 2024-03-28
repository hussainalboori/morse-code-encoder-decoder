package checkdata

import "strings"

func IsMorseCode(input string) bool {
	// Define a set of Morse code characters
	morseChars := []string{".", "-", " ", "/"}

	// Check if the input string only contains Morse code characters
	for _, char := range input {
		if !strings.ContainsAny(string(char), strings.Join(morseChars, "")) {
			return false
		}
	}

	return true
}
