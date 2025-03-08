package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

func DeadlineMiddleware(next http.Handler, timeout time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", DeadlineMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(11 * time.Second)
		select {
		case <-r.Context().Done():
			w.WriteHeader(http.StatusRequestTimeout)
			json.NewEncoder(w).Encode(map[string]string{"message": "Request cancelled"})
			slog.Info("Request cancelled")
			return
		default:
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
		}
	}), 10*time.Second))
	slog.Info("Starting server on port 8080...", "port", 8080)
	http.ListenAndServe(":8080", mux)
}
