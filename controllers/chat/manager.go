package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

var Manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (Manager *ClientManager) Start() {
	for {
		select {
		case conn := <-Manager.register:
			Manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			Manager.Send(jsonMessage, conn)
		case conn := <-Manager.unregister:
			if _, ok := Manager.clients[conn]; ok {
				fmt.Println(conn.id + " - unregister")

				close(conn.send)
				delete(Manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				Manager.Send(jsonMessage, conn)
			}
		case message := <-Manager.broadcast:
			for conn := range Manager.clients {
				select {
				case conn.send <- message:
				}
			}
		}
	}
}

func (Manager *ClientManager) Send(message []byte, ignore *Client) {
	for conn := range Manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}
func NewClient(res http.ResponseWriter, req *http.Request) {
	fmt.Println("new client ws connting...")
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		fmt.Println("new client error: " + err.Error())
		http.NotFound(res, req)
		return
	}
	id, _ := uuid.NewV4()
	client := &Client{id: id.String(), socket: conn, send: make(chan []byte)}

	Manager.register <- client
	fmt.Println(id.String() + " -success register new client ws conntion")

	go client.Read()
	go client.Write()
}
