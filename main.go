package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileHandler := http.FileServer(http.Dir("."))
	http.Handle("GET /", fileHandler)
	log.Println("creating server at:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
