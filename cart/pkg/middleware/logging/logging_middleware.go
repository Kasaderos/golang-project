package logging

import (
	"net/http"
)

// WithHTTPLoggingMiddleware
func WithHTTPLoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
