package main

import (
	"log"
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type:", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))

}

func main() {
	mux := http.NewServeMux()

	fileHandler := http.FileServer(http.Dir("."))
	mux.Handle("/app/", http.StripPrefix("/app", fileHandler))
	imageHandler := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets", http.StripPrefix("/assets", imageHandler))

	mux.HandleFunc("GET /healthz", healthz)

	log.Println("creating server at:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
