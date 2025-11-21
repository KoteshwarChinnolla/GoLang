package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimit(h http.Handler) http.Handler {
	limiter := rate.NewLimiter(1, 3)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		h.ServeHTTP(w, r)
	})
}
