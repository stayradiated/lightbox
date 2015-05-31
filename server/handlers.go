package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/stayradiated/lightbox/server/db"

	"github.com/gorilla/mux"
)

type Handlers struct {
	DB *db.DB
}

// ReadShows returns a list of show
func (h Handlers) ReadShows(w http.ResponseWriter, r *http.Request) {

	filter := r.FormValue("filter")
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		limit = 24
	}
	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		offset = 0
	}

	shows, err := h.DB.Shows(filter, limit, offset)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	printJson(w, shows)
}

// ReadShow returns a single show
func (h Handlers) ReadShow(w http.ResponseWriter, r *http.Request) {

	showID, err := strconv.Atoi(mux.Vars(r)["show"])
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	show, err := h.DB.Show(showID)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	printJson(w, show)
}

// ReadCategory returns a list of shows in a category
func (h Handlers) ReadCategory(w http.ResponseWriter, r *http.Request) {

	categoryID, err := strconv.Atoi(mux.Vars(r)["category"])
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	shows, err := h.DB.CategoryShows(categoryID)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	printJson(w, shows)
}

// ReadSeason returns a single season
func (h Handlers) ReadSeason(w http.ResponseWriter, r *http.Request) {

	seasonID, err := strconv.Atoi(mux.Vars(r)["season"])
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	season, err := h.DB.Season(seasonID)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	printJson(w, season)
}

// ReadEpisode returns a single episode
func (h Handlers) ReadEpisode(w http.ResponseWriter, r *http.Request) {

	episodeID, err := strconv.Atoi(mux.Vars(r)["episode"])
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	episode, err := h.DB.Episode(episodeID)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	printJson(w, episode)
}

// printJson
func printJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(obj)
}
