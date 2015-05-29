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
			"ReadSeriesList",
			"GET", "/series", h.ReadSeries,
		},
		Route{
			"ReadSeriesWithID",
			"GET", "/series/{series:[0-9]+}", h.ReadSeriesWithID,
		},
		Route{
			"ReadSeason",
			"GET", "/season/{season:[0-9]+}", h.ReadSeason,
		},
	}
}
