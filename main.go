package main

import (
	"fmt"
	"strings"
)

var morseCode = map[string]string{
	".-":   "A", "-...": "B", "-.-.": "C", "-..":  "D", ".":    "E",
	"..-.": "F", "--.":  "G", "....": "H", "..":   "I", ".---": "J",
	"-.-":  "K", ".-..": "L", "--":   "M", "-.":   "N", "---":  "O",
	".--.": "P", "--.-": "Q", ".-.":  "R", "...":  "S", "-":    "T",
	"..-":  "U", "...-": "V", ".--":  "W", "-..-": "X", "-.--": "Y",
	"--..": "Z", ".----": "1", "..---": "2", "...--": "3", "....-": "4",
	".....": "5", "-....": "6", "--...": "7", "---..": "8", "----.": "9",
	"-----": "0", "/": " ",
}

func main(){
	var morseCodeInpout string = "--. --. / --. --."
	// fmt.Println("enter the massge you want to encode or decode:")
	// fmt.Scan(&morseCodeInpout)
	messge := decodeMorseCode(morseCodeInpout)
	fmt.Println(messge)
	//fmt.Println(morseCodeInpout)
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