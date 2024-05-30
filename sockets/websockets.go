package sockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Impossible d'ouvrir une connexion websocket", http.StatusBadRequest)
		fmt.Printf("Erreur lors de la mise a niveau WebSocket : %v\n", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Erreur lors de la lecture du message : %v", err)
			break
		}
		// Echo the message back
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Printf("Erreur lors de l'ecriture du message : %v", err)
			break
		}
	}
}
