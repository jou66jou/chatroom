package main

import (
	"fmt"
	cs "go-chat-room/controllers"
	"go-chat-room/routers"
	"log"
	"net/http"
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
