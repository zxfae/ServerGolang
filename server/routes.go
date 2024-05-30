package server

import (
	"net/http"
	"real-time-forum/handling"
)

func InitRoutes() http.Handler{
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("frontend"))))

	mux.HandleFunc("/", handling.Handlemain)
	// http.HandleFunc("/", handling.HandleRegister)
	// http.HandleFunc("/register", handling.HandleRegister)
	// http.HandleFunc("/ws", sockets.WebSocketHandler)
	return secureHeaders(mux)
}
