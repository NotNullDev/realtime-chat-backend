package ws

import (
	"encoding/json"
	"log"

	t "github.com/NotNullDev/realtime-chat-backend/types"
	"github.com/gorilla/websocket"
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

		// messageType, messageContent, err := conn.ReadMessage()

		// if err != nil {
		// 	log.Println(err)
		// }

		// log.Printf("Received message of type [%d] with content: [%s]", messageType, messageContent)
	}
}

func handleMessage(conn *websocket.Conn) {
	// var msg t.ChatMessage

	for {
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
}
