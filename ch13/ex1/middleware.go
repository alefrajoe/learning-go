package main

import (
	"log"
	"net/http"
)

func loggingIncomingIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		log.Printf("Incoming request from %s", ip)
		next.ServeHTTP(w, r)
	})
}
	