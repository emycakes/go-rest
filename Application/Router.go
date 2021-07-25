package Application

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	mux *mux.Router
}

func (router Router) Listen(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router.mux))
}

func (router Router) RegisterRoutes(routes RouteCollection) {
	for _, route := range routes {
		router.mux.HandleFunc(route.Pattern, createHandler(route)).Methods(route.Method)
	}
}

func createHandler(route Route) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		route.Handler(Request{Pointer: request}, Response{Writer: responseWriter})
	}
}

func NewRouter() Router {
	return Router{
		mux: mux.NewRouter().StrictSlash(true),
	}
}