package http

import (
	"log"
	"net/http"

	"github.com/NotNullDev/realtime-chat-backend/ws"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // for development only
	},
}

func WsHttpEndpointHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("New connection started")
	websocketConnection, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Connection upgraded.")

	go ws.HandleWebsocketClient(websocketConnection)
}

func chatMessageModelSample(w http.ResponseWriter, r *http.Request) {

}
