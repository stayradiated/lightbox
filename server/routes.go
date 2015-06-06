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
			"GET", "/shows.json", h.ReadShows,
		},
		Route{
			"ReadShow",
			"GET", "/shows/{show:[0-9]+}.json", h.ReadShow,
		},
		Route{
			"ReadSeason",
			"GET", "/seasons/{season:[0-9]+}.json", h.ReadSeason,
		},
		Route{
			"ReadCategory",
			"GET", "/categories/{category:[0-9]+}.json", h.ReadCategory,
		},
		Route{
			"ReadCategories",
			"GET", "/categories.json", h.ReadCategories,
		},
		Route{
			"ReadEpisode",
			"GET", "/episodes/{episode:[0-9]+}.json", h.ReadEpisode,
		},
		Route{
			"ReadLists",
			"GET", "/lists.json", h.ReadLists,
		},
	}
}
