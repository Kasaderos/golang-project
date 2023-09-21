package controller_http

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

func MiddlewareRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("err: %v, stacktrace %s\n", err, debug.Stack()) // May be log this error? Send to sentry?

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)
	})
}
