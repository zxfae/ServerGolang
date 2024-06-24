package server

import (
	"net/http"
	"time"
)

// ServerParameters configures the server parameters
func ServerParameters(handler http.Handler, Request int) *http.Server {
	return &http.Server{
		Addr:              "127.0.0.1:8080",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Handler:           MiddlewareRateLimiting(handler, Request),
	}
}
