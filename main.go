package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var morseCode = map[string]string{
	".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E",
	"..-.": "F", "--.": "G", "....": "H", "..": "I", ".---": "J",
	"-.-": "K", ".-..": "L", "--": "M", "-.": "N", "---": "O",
	".--.": "P", "--.-": "Q", ".-.": "R", "...": "S", "-": "T",
	"..-": "U", "...-": "V", ".--": "W", "-..-": "X", "-.--": "Y",
	"--..": "Z", ".----": "1", "..---": "2", "...--": "3", "....-": "4",
	".....": "5", "-....": "6", "--...": "7", "---..": "8", "----.": "9",
	"-----": "0", "/": " ",
}

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

func isMorseCode(input string) bool {
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

func main() {
	//result := ""
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("entar the line you wnat to translate: ")
	inpout, _ := reader.ReadString('\n')
	inpout = strings.TrimSpace(inpout)
	text := isMorseCode(inpout)
	if text == true {
		result := decodeMorseCode(inpout)
		fmt.Println(result)
	} else {
		result := encodeMorseCode(inpout)
		fmt.Println(result)
	}
}

func decodeMorseCode(code string) string {
	words := strings.Split(code, " / ")
	decodedMessage := ""

	for _, word := range words {
		symbols := strings.Split(word, " ")
		for _, symbol := range symbols {
			decodedSymbol, exists := morseCode[symbol]
			if exists {
				decodedMessage += decodedSymbol
			}
		}
		decodedMessage += " "
	}

	return decodedMessage
}

func encodeMorseCode(text string) string {
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
