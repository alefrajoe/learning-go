package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log(r.Context(), Debug, "Hello, World!")
	})))
	log.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", mux)
}
