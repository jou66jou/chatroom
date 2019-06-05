package controllers

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
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
		// fmt.Println(c.id + " : " + string(message))
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
