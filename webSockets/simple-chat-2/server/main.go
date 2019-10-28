package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configuring the upgrader
var upgrader = websocket.Upgrader{}

// Message object contains data from client
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// Creating a file server
	fs := http.FileServer(http.Dir("../client"))
	http.Handle("/", fs)

	// Configuring websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8000 and log if any error encountered
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Error occured: ", err)
	}
}

// Function to handle incomming websocket connection
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrading initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error getting connection: ", err)
	}
	defer ws.Close()

	// Registering new client
	clients[ws] = true

	for {
		var msg Message
		// Reading new message as JSON and mapping it to the Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error: Client disconnected: %v", err)
			delete(clients, ws)
			break
		}
		// Sending the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grabing message from broadcast server
		msg := <-broadcast
		// Sending the message to every clients who are currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error: Client might be disconnected: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
