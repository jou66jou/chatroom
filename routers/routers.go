package routers

import (
	"fmt"
	"net/http"

	controllers "github.com/jou66jou/go-chat-room/controllers/chat"

	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	// fmt.Println("HTTP Method list:")
	fmt.Println("Websocket : /chatroom - call the websocket into chatroom")
	register("", "/chatroom", controllers.NewClient, nil)
	fmt.Println("")

}

func Routers() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		var r *mux.Route
		if route.Method != "" {
			r = router.Methods(route.Method).
				Path(route.Pattern)
		} else {
			r = router.Path(route.Pattern)
		}

		if route.Middleware != nil { // JWT valid
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
