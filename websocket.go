package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make([]*websocket.Conn, 0)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func hendleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error Connection", err)
		return
	}

	clients = append(clients, conn)

	fmt.Println("New client connected")

	//read message for client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error Message", err)
			break
		}

		for _, client := range clients {

			if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("Error sending message:", err)
			}
		}
	}

	for i, client := range clients {
		if client == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}

}
