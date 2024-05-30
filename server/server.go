package server

import (
	"net/http"
	"time"
	"fmt"
)

/*
Server parameters,
DOS, servToClient && clientToServ communication,
Datas length && more
*/
func ServerParameters(handler http.Handler) *http.Server {
	return &http.Server{
		Addr: "127.0.0.1:8080",
		// DOS protections
		ReadTimeout: 10 * time.Second,
		// Header protection, if client send datas to server
		ReadHeaderTimeout: 5 * time.Second,
		// Datas send serv to client
		WriteTimeout: 10 * time.Second,
		// Keep_Alives protection
		IdleTimeout: 120 * time.Second,
		// Datas send by client to serv, protected less 1MB
		// 1000000000000000000000 bits == 1 MB
		// Return 431 statusCode
		MaxHeaderBytes: 1 << 20,
		Handler: secureHeaders(handler),
	}
}

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}
