package ws

import (
	"encoding/json"
	t "github.com/NotNullDev/realtime-chat-backend/types"
	"github.com/gorilla/websocket"
	"log"
)

var ChatServer = t.WsServer{
	RegisterChannel:   make(chan *t.ChatUser),
	UnregisterChannel: make(chan *t.ChatUser),
}
var activeChannels = []t.ChatChannel{}
var activeUsers = []t.ChatUser{}

func HandleWebsocketClient(conn *websocket.Conn) {
	for {
		handleMessage(conn)
	}
}

func handleMessage(conn *websocket.Conn) {
	// var msg t.ChatMessage
	msgType, msg, err := conn.ReadMessage()

	if err != nil {
		log.Printf("ERROR: could not read message. Error: [%s]", err.Error())
	}

	log.Printf("Received message wity type %d and content: %v", msgType, string(msg))

	log.Printf("Trying to translate message...")

	var translatedMessage t.ChatMessage

	err = json.Unmarshal(msg, &translatedMessage)

	if err != nil {
		log.Println(err.Error())
	}
}
