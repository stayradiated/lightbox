package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func GetRoutes(h *Handlers) Routes {
	return Routes{
		Route{
			"ReadShows",
			"GET", "/shows", h.ReadShows,
		},
		Route{
			"ReadShow",
			"GET", "/shows/{show:[0-9]+}", h.ReadShow,
		},
		Route{
			"ReadSeason",
			"GET", "/seasons/{season:[0-9]+}", h.ReadSeason,
		},
	}
}
