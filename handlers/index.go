package handlers

import (
	"html/template"
	"log"
	decode "morse-code/Decode"
	encoder "morse-code/Encoder"
	checkdata "morse-code/checkData"
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
	Data := ""
	if checkdata.IsMorseCode(text) == true {
		Data = decode.DecodeMorseCode(text)
	} else {
		Data = encoder.EncodeMorseCode(text)
	}
	
	err = template.Execute(w, Data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		errTemplate, err := template.ParseFiles("templates/err500.html")
		if err != nil {
			log.Println("Error parsing error template:", err)
			return
		}
		errTemplate.Execute(w, nil)
		return
	}
}
