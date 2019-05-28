package routers

import (
	"fmt"
	"net/http"
	"net/http/pprof"

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
	register("GET", "/debug/pprof/", pprof.Index, nil)
	register("GET", "/debug/pprof/cmdline", pprof.Cmdline, nil)
	register("GET", "/debug/pprof/profile", pprof.Profile, nil)
	register("GET", "/debug/pprof/symbol", pprof.Symbol, nil)
	register("GET", "/debug/pprof/trace", pprof.Trace, nil)
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
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
