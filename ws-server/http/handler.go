package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type WsAuthRequest struct {
	UserId string `json:"userId"`
	RoomId int64  `json:"roomId"`
	Secret string `json:"secret"`
}

func isAuthorized(data WsAuthRequest) bool {
	dataAsString, err := json.Marshal(data)

	if err != nil {
		log.Printf("Something went wring during marshaling data: [%v] [%s]", data, err)
	}

	resp, err := http.Post("http://localhost:3000/api/verifyAuthorization", "application/json", bytes.NewBuffer(dataAsString))

	if err != nil {
		log.Printf("Could not send http request: [%s].", err)
	}

	return resp.StatusCode == 200
}

func WsHttpEndpointHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("New connection started")
	bodyAsString, err := io.ReadAll(r.Body)

	if err != nil {
		errMsg := fmt.Sprintf("ERROR: could not read body: [%s]", err.Error())

		log.Printf(errMsg)
		w.WriteHeader(500)
		w.Write([]byte(errMsg))
		return
	}

	var body WsAuthRequest
	err = json.Unmarshal(bodyAsString, &body)

	if err != nil {
		w.WriteHeader(422)
		errMsg := fmt.Sprintf("Could not parse body [%s]", bodyAsString)
		w.Write([]byte(errMsg))
		return
	}

	if isAuthorized(body) != true {
		log.Printf("Authorization no granted for request with data: [%v]", body)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	log.Printf("Auhorization granted for request with data: [%v]", body)

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
