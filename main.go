package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Printf("Starting HTTP Web Server")
	log.Fatal(http.ListenAndServe(":80", nil))
}
