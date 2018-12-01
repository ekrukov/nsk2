package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Start application")

	port := os.Getenv("PORT")

	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", helthz)
	http.HandleFunc("/readyz", ready)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server is stopped with error: %v", err)
	}

	log.Print("Stop application")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func helthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
