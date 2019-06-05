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
	ip := "192.168.50.181"
	port := "8080"
	count := 10000
	fmt.Println("Server ip (default " + ip + "):")
	SwitchScanf(&ip)

	fmt.Println("Server port (default " + port + "):")
	SwitchScanf(&port)

	fmt.Println("client count (default 10000):")
	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		count = 10000
	}
	var addr = flag.String("addr", ip+":"+port, "http service address")
	for i := 0; i < count; i++ {
		u := url.URL{Scheme: "ws", Host: *addr, Path: "/chatroom"}
		var dialer *websocket.Dialer

		conn, _, err := dialer.Dial(u.String(), nil)
		if err != nil {
<<<<<<< HEAD
			fmt.Println(err)
=======
			fmt.Println("conn err: " + err.Error())
>>>>>>> 494a0705004015722c36b1954af2fc2ff126f94b
			break
		}
		wg.Add(1)
		// go timeWriter(i, conn)
		go wsRead(i, conn)
<<<<<<< HEAD
		fmt.Println("create client: %d", i)
		time.Sleep(1 * time.Millisecond)

	}
	fmt.Println("done")
=======
		fmt.Printf("create client: %d\n", i)
		time.Sleep(1 * time.Millisecond)

	}
	fmt.Printf("create client done: %d\n", count)

>>>>>>> 494a0705004015722c36b1954af2fc2ff126f94b
	wg.Wait()
}

func wsRead(i int, conn *websocket.Conn) {
	conn.SetReadDeadline(time.Now().Add(100 * time.Minute))
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read ", "No."+strconv.Itoa(i)+" : ", err)
			wg.Done()
			return
		}

		// fmt.Printf("received: %s\n", message)
	}
	wg.Done()

}
func timeWriter(i int, conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 10)
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
