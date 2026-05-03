package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer conn.Close()

	log.Printf("Client %v connected.", r.RemoteAddr)

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			break
		}
	}
	log.Printf("Client %v disconnected.", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("WS server running on :4000")

	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatalln(err)
	}
}
