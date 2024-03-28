package main

import (
	"log"
	"morse-code/handlers"
	"net/http"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", handlers.IndexHandler)
	log.Printf("Connect to our website through http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
