package main

import (
	"flag"
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var wg sync.WaitGroup

func main() {
	var port string = "8080"
	fmt.Println("Server port (default 8080):")
	SwitchScanf(&port)
	var addr = flag.String("addr", "localhost:"+port, "http service address")
	for i := 0; i < 5; i++ {
		u := url.URL{Scheme: "ws", Host: *addr, Path: "/chatroom"}
		var dialer *websocket.Dialer

		conn, _, err := dialer.Dial(u.String(), nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		wg.Add(2)
		go timeWriter(i, conn)
		go wsRead(i, conn)
		time.Sleep(time.Second * 2)

	}
	wg.Wait()
}

func wsRead(i int, conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read ", "No."+strconv.Itoa(i)+" : ", err)
			wg.Done()
			return
		}

		fmt.Printf("received: %s\n", message)
	}
	wg.Done()

}
func timeWriter(i int, conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 2)
		conn.WriteMessage(websocket.TextMessage, []byte("No."+strconv.Itoa(i)+" : "+time.Now().Format("2006-01-02 15:04:05")))
	}
	wg.Done()

}

func SwitchScanf(v *string) {
	var s string
	fmt.Scanln(&s)
	if s != "" {
		*v = s
	}
}