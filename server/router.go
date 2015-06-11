package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(handlers *Handlers) *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	// load routes
	for _, route := range GetRoutes(handlers) {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		api.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// serve client files (html, css, etc)
	client := http.FileServer(http.Dir("../gh-pages"))
	router.PathPrefix("/").Handler(client)

	return router
}
