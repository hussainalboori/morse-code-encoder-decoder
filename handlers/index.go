package handlers

import (
	"html/template"
	"log"
	"morse-code/code"
	"net/http"
)

// EncodeText encodes the given text into Morse code and sends it back to the front end.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Get the text from the front end
	text := r.FormValue("text")
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	morseCode := code.EncodeMorseCode(text)
	template.Execute(w, morseCode)
	log.Println(morseCode)

	// TODO: Encode the text into Morse code

	// TODO: Send the encoded Morse code back to the front end
}
