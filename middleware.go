package main

import (
	"log"
	"net/http"
	"time"
)

type Middleware struct{}

func (m Middleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v", r.Method, r.URL, t2.Sub(t1))
	})
}
