package server

import (
	"net/http"

	"golang.org/x/time/rate"
)

// MiddlewareRateLimiting limits the number of requests per second
func MiddlewareRateLimiting(next http.Handler, Request int) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(Request), Request)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
