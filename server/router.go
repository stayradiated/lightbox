package main

import (
	"compress/gzip"
	"net/http"

	"github.com/carbocation/interpose"
	"github.com/carbocation/interpose/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(handlers *Handlers) *interpose.Middleware {

	middle := interpose.New()

	middle.Use(middleware.NegroniGzip(gzip.DefaultCompression))

	router := mux.NewRouter()
	middle.UseHandler(router)

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

	return middle
}
