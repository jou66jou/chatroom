package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	cs "github.com/jou66jou/go-chat-room/controllers/chat"
	"github.com/jou66jou/go-chat-room/routers"
)

func main() {
	var port string = "8080"
	fmt.Println("Http listen port (default 8080):")
	SwitchScanf(&port)
	// starts an chatroom service
	go cs.Manager.Start()
	fmt.Println("HTTP Listening on " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, routers.Routers()))

}
func SwitchScanf(v *string) {
	var s string
	fmt.Scanln(&s)
	if s != "" {
		*v = s
	}
}
