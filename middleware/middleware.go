// middleware/logging.go

package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs each incoming request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request method and URL path
		log.Printf("[%s] %s %s", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
