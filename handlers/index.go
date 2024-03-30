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
	if r.URL.Path != "/" {
		log.Printf("%s tried to access a non-existent page: %s\n", r.RemoteAddr, r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		template, err := template.ParseFiles("templates/err404.html")
		if err != nil {
			log.Println("Error parsing 404 template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		template.Execute(w, nil)
		return
	}
	// Get the text from the front end
	text := r.FormValue("text")
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("Reciving Text:", text)
	Data := ""
	if checkdata.IsMorseCode(text) == true {
		log.Println("Decoding Morse Code")
		Data = decode.DecodeMorseCode(text)
	} else {
		Data = encoder.EncodeMorseCode(text)
		log.Println("Encoding Morse Code")
	}

	err = template.Execute(w, Data)
	Data = ""
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

	if r.Method == "POST" {
		log.Printf("%s made a POST request to the root path\n", r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)
		template.ExecuteTemplate(w, "err400.html", nil)
	}

}
