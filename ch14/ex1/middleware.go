package main

import (
	"context"
	"net/http"
	"time"
)

// DeadlineMiddleware is a middleware that cancels the request after a timeout
// and returns a 408 Request Timeout status code.
//
// It uses the context.WithTimeout function to create a new context with a timeout.
// The new context is then used to serve the request.
//
// If the request is cancelled, the middleware returns a 408 Request Timeout status code.
// Otherwise, the middleware serves the request and returns the response from the next handler.
func DeadlineMiddleware(next http.Handler, timeout time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
