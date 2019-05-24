package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

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
