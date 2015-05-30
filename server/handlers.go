package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handlers struct {
	DB *sql.DB
}

// ReadShows returns a list of show
func (h Handlers) ReadShows(w http.ResponseWriter, r *http.Request) {

	filter := "%" + r.FormValue("filter") + "%"

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if limit == 0 || err != nil {
		limit = 24
	}
	// if limit > 50 {
	// 	limit = 50
	// }

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if offset == 0 || err != nil {
		offset = 0
	}

	rows, err := h.DB.Query(`select
		id, name, poster, first_aired
		from shows
		where name like (?)
		order by first_aired desc
		limit ? offset ?`, filter, limit, offset)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	shows := make([]Show, 0)

	for rows.Next() {
		show := Show{}
		err := rows.Scan(&show.ID, &show.Name, &show.Poster, &show.FirstAired)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		shows = append(shows, show)
	}

	printJson(w, shows)
}

// ReadShow returns a single show
func (h Handlers) ReadShow(w http.ResponseWriter, r *http.Request) {

	showID := mux.Vars(r)["show"]
	show := Show{}

	err := h.DB.QueryRow(`select
		id,
		name,
		overview,
		rating,
		rating_count,
		actors,
		poster,
		banner,
		fanart,
		content_rating,
		first_aired,
		runtime,
		imdb
	from shows
	where id = (?)`, showID).Scan(
		&show.ID,
		&show.Name,
		&show.Overview,
		&show.Rating,
		&show.RatingCount,
		&show.Actors,
		&show.Poster,
		&show.Banner,
		&show.Fanart,
		&show.ContentRating,
		&show.FirstAired,
		&show.Runtime,
		&show.IMDB,
	)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	rows, err := h.DB.Query(`
		select categories.name
		from categories, show_categories, shows
		where 
			categories.id = show_categories.category_id and
			shows.id = show_categories.show_id and
			shows.id = ?
	`, showID)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	categories := make([]string, 0)

	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			fmt.Fprintln(w, err)
			return
		}
		categories = append(categories, category)
	}

	show.Categories = categories

	rows, err = h.DB.Query(`select
		id, show_id, number, banner
		from seasons
		where show_id = (?)`, showID)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	seasons := make([]Season, 0)

	for rows.Next() {
		var season Season
		if err := rows.Scan(&season.ID, &season.ShowID, &season.Number, &season.Banner); err != nil {
			fmt.Fprintln(w, err)
			return
		}
		seasons = append(seasons, season)
	}

	show.Seasons = seasons

	printJson(w, show)
}

// ReadSeason reads a list of show
func (h Handlers) ReadSeason(w http.ResponseWriter, r *http.Request) {

	seasonID := mux.Vars(r)["season"]
	season := Season{}

	err := h.DB.QueryRow(`select
		id, show_id, number, banner
		from seasons
		where id = (?)`, seasonID).Scan(&season.ID, &season.ShowID, &season.Number, &season.Banner)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	rows, err := h.DB.Query(`select
		id,
		episode_number,
		episode_name,
		overview,
		first_aired,
		filename,
		rating,
		rating_count,
		director,
		writer,
		guest_stars,
		imdb
	from episodes where season_id = (?)`, seasonID)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	episodes := make([]Episode, 0)

	for rows.Next() {
		episode := Episode{}
		err = rows.Scan(
			&episode.ID,
			&episode.Number,
			&episode.Name,
			&episode.Overview,
			&episode.FirstAired,
			&episode.Image,
			&episode.Rating,
			&episode.RatingCount,
			&episode.Director,
			&episode.Writer,
			&episode.GuestStars,
			&episode.IMDB,
		)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		episodes = append(episodes, episode)
	}

	season.Episodes = episodes

	printJson(w, season)
}

// printJson
func printJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(obj)
}
