// https://golangdocs.com/golang-gorilla-websockets

// server.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

// socketHandler
func socketHandler(w http.ResponseWriter, r *http.Request) {

	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}

		log.Printf("Server Recieved: %s", message)

		// Convert byte[] to string
		myMsg := "Pong: " + string(message)

		// Convert string to byte
		byteMsg := []byte(myMsg)

		err = conn.WriteMessage(messageType, byteMsg)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}

		// Causes Hijack Connection
		// fmt.Fprintf(w, "Socket ", message)
	}
}

// home
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page "+r.Method+" from Host Address "+r.Host)

	//
	fmt.Println("")
	fmt.Fprintf(w, "   https://pkg.go.dev/net/http#Request")
}

func main() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
