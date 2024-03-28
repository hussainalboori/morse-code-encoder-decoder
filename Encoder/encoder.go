package encoder

import "strings"

var morseDict = map[string]string{
	"A": ".-", "B": "-...", "C": "-.-.", "D": "-..", "E": ".",
	"F": "..-.", "G": "--.", "H": "....", "I": "..", "J": ".---",
	"K": "-.-", "L": ".-..", "M": "--", "N": "-.", "O": "---",
	"P": ".--.", "Q": "--.-", "R": ".-.", "S": "...", "T": "-",
	"U": "..-", "V": "...-", "W": ".--", "X": "-..-", "Y": "-.--",
	"Z": "--..", "1": ".----", "2": "..---", "3": "...--", "4": "....-",
	"5": ".....", "6": "-....", "7": "--...", "8": "---..", "9": "----.",
	"0": "-----", " ": "/",
}

func EncodeMorseCode(text string) string {
	text = strings.ToUpper(text)
	encodedMessage := ""

	for _, char := range text {
		encodedSymbol, exists := morseDict[string(char)]
		if exists {
			encodedMessage += encodedSymbol + " "
		}
	}

	return encodedMessage
}
