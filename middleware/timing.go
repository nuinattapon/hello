package middleware

import (
	"log"
	"net/http"
	"time"
)

// TimingMiddleware adds timing information to HTTP handlers
func TimingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler
		next(w, r)

		// Calculate and log the elapsed time
		elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0
		log.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
	}
}
