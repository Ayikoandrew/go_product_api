package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"main.go/handlers"
)

func LogReport(l *log.Logger, hd *handlers.Product) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Printf("Method: %s, URL: %s, %s\n", r.Method, r.RequestURI, r.RemoteAddr)
			next.ServeHTTP(w, r)
			l.Printf("Completed in %v\n", time.Since(start))
		})
	}
}
