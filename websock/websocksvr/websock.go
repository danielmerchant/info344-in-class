package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

/*
TODO: Implement the code in this file, according to the comments.
If you haven't yet read the assigned reading, now would be a
good time to do so:
- Read the Overview section of the Gorilla WebSockets package
https://godoc.org/github.com/gorilla/websocket
- Read the Writing WebSocket Client Application
https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_client_applications
*/

//WebSocketsHandler is a handler for WebSocket upgrade requests
type WebSocketsHandler struct {
	notifier *Notifier
	upgrader *websocket.Upgrader
}

//NewWebSocketsHandler constructs a new WebSocketsHandler
func NewWebSocketsHandler(notifer *Notifier) *WebSocketsHandler {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	return &WebSocketsHandler{notifier: notifer, upgrader: &upgrader}
}

//ServeHTTP implements the http.Handler interface for the WebSocketsHandler
func (wsh *WebSocketsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("received websocket upgrade request")
	conn, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("error connecting: %v", err)
	}
	wsh.notifier.AddClient(conn)
	//TODO: Upgrade the connection to a WebSocket, and add the
	//new websock.Conn to the Notifier. See
	//https://godoc.org/github.com/gorilla/websocket#hdr-Overview
}

//Notifier is an object that handles WebSocket notifications
type Notifier struct {
	clients []*websocket.Conn
	eventQ  chan []byte
	mu      sync.RWMutex
	//TODO: add a mutex or other channels to
	//protect the `clients` slice from concurrent use.
}

//NewNotifier constructs a new Notifier
func NewNotifier() *Notifier {
	n := &Notifier{}
	n.eventQ = make(chan []byte, 100)
	go n.start()
	return n
}

//AddClient adds a new client to the Notifier
func (n *Notifier) AddClient(client *websocket.Conn) {
	log.Println("adding new WebSockets client")
	n.mu.Lock()
	n.clients = append(n.clients, client)
	n.mu.Unlock()
	//TODO: add the client to the `clients` slice
	//but since this can be called from multiple
	//goroutines, make sure you protect the `clients`
	//slice while you add a new connection to it!

	//also process incoming control messages from
	//the client, as described in this section of the docs:
	//https://godoc.org/github.com/gorilla/websocket#hdr-Control_Messages
	nClients := []*websocket.Conn{}
	for {
		if _, _, err := client.NextReader(); err != nil {
			client.Close()
			n.mu.Lock()
			for _, c := range n.clients {
				if c != client {
					nClients = append(nClients, c)
				}
			}
			n.clients = nClients
			n.mu.Unlock()
			break
		}
	}
}

//Notify broadcasts the event to all WebSocket clients
func (n *Notifier) Notify(event []byte) {
	log.Printf("adding event to the queue")
	n.mu.Lock()
	n.eventQ <- event
	n.mu.Unlock()
	//TODO: add `event` to the `n.eventQ`
}

//start starts the notification loop
func (n *Notifier) start() {
	log.Println("starting notifier loop")
	for {
		event := <-n.eventQ
		n.mu.Lock()
		for _, c := range n.clients {
			if err := c.WriteMessage(websocket.TextMessage, event); err != nil {
				log.Printf("error sending message: %v", err)
				return
			}
		}
		n.mu.Unlock()
	}
	//TODO: start a never-ending loop that reads
	//new events out of the `n.eventQ` and broadcasts
	//them to all WebSocket clients.
	//If you use additional channels instead of a mutex
	//to protext the `clients` slice, also process those
	//channels here using a non-blocking `select` statement
}
