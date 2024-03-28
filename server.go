package main

import (
	"log"
	"morse-code/handlers"
	"net/http"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", handlers.IndexHandler)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	log.Printf("Connect to our website through http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
