package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

const port = ":8080"

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error while upgrading connection: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("client connected")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error while reading message: ", err)
			break
		}

		fmt.Printf("received message: %s\n", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("error while writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)

	fmt.Println("websocket server started on " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("error while starting server: ", err)
	}
}
