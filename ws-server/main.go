package main

import (
	"fmt"
	"github.com/NotNullDev/realtime-chat-backend/db"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/NotNullDev/realtime-chat-backend/mylog"
	"github.com/NotNullDev/realtime-chat-backend/ws"

	h "github.com/NotNullDev/realtime-chat-backend/http"
	c "github.com/NotNullDev/realtime-chat-backend/myconst"
)

func main() {
	logFile, err := os.Create("log.log")

	if err != nil {
		panic(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Printf("ERROR: %s\n", err)
		}
	}(logFile)

	mw := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(mw)

	db.Init()
	http.HandleFunc("/ws", h.WsHttpEndpointHandler)
	http.HandleFunc("/joinToRoom", func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)

		if err != nil {
			log.Printf("ERROR: could not read body")
			return
		}

		log.Printf("Received body: [%s]", string(body))
		log.Printf("Ip address: %v", request.RemoteAddr)

		log.Printf("Headers:\n")
		for name, values := range request.Header {

			log.Printf("Header: %s. Values:\n", name)

			for idx, val := range values {
				log.Printf("%d. %s", idx, val)
			}
		}

		writer.Header().Add("OMG", "XD")
		writer.WriteHeader(200)
		write, err := writer.Write([]byte("ok"))

		if err != nil {
			log.Printf("ERROR: could not write response. error: [%s]", err.Error())
		} else {
			log.Printf("Bytes written: %d", write)
		}
	})

	log.Printf("Started server on: %s", c.HTTP_SERVER_LISTEN_ADDRESS)

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go handleHttpConnections(wg)

	wg.Add(1)
	go handleWsRegistration(wg)

	wg.Wait()
}

func handleWsRegistration(wg *sync.WaitGroup) {
	register := ws.ChatServer.RegisterChannel
	unregister := ws.ChatServer.RegisterChannel

	select {
	case client := <-register:
		fmt.Printf("New client registred: [%s]", client.Id)
	case client := <-unregister:
		fmt.Printf("Client unregistered: [%s]", client.Id)
	}

	wg.Done()
}

func handleHttpConnections(wg *sync.WaitGroup) {
	err := http.ListenAndServe(c.HTTP_SERVER_LISTEN_ADDRESS, nil)
	if err != nil {
		mylog.Error(err)
	}

	wg.Done()
}
