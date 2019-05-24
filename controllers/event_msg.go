package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

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
					fmt.Println(conn.id + " - broadcast send")

					// default:
					// 	fmt.Println(conn.id + " - broadcast close")

					// 	// close(conn.send)
					// 	// delete(Manager.clients, conn)
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

func (c *Client) Read() {
	defer func() {
		Manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			Manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		fmt.Println(c.id + " : " + string(message))
		Manager.broadcast <- jsonMessage
	}
}

func (c *Client) Write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
