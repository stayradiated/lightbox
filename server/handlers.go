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

// ReadSeries reads a list of series
func (h Handlers) ReadSeries(w http.ResponseWriter, r *http.Request) {

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
		id, series_name, poster, first_aired
		from series
		where series_name like (?)
		order by series_name
		limit ? offset ?`, filter, limit, offset)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	seriesList := make([]Series, 0)

	for rows.Next() {
		series := Series{}
		err := rows.Scan(&series.ID, &series.Name, &series.Poster, &series.FirstAired)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		seriesList = append(seriesList, series)
	}

	printJson(w, seriesList)
}

// ReadSeries reads a list of series
func (h Handlers) ReadSeriesWithID(w http.ResponseWriter, r *http.Request) {

	seriesID := mux.Vars(r)["series"]
	series := Series{}

	err := h.DB.QueryRow(`select
		id,
		series_name,
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
	from series
	where id = (?)`, seriesID).Scan(
		&series.ID,
		&series.Name,
		&series.Overview,
		&series.Rating,
		&series.RatingCount,
		&series.Actors,
		&series.Poster,
		&series.Banner,
		&series.Fanart,
		&series.ContentRating,
		&series.FirstAired,
		&series.Runtime,
		&series.IMDB,
	)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	rows, err := h.DB.Query(`
		select categories.name
		from categories, series_categories, series
		where 
			categories.id = series_categories.category_id and
			series.id = series_categories.series_id and
			series.id = ?
	`, seriesID)

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

	series.Categories = categories

	rows, err = h.DB.Query(`select
		id, series_id, number, banner
		from seasons
		where series_id = (?)`, seriesID)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	seasons := make([]Season, 0)

	for rows.Next() {
		var season Season
		if err := rows.Scan(&season.ID, &season.SeriesID, &season.Number, &season.Banner); err != nil {
			fmt.Fprintln(w, err)
			return
		}
		seasons = append(seasons, season)
	}

	series.Seasons = seasons

	printJson(w, series)
}

// ReadSeason reads a list of series
func (h Handlers) ReadSeason(w http.ResponseWriter, r *http.Request) {

	seasonID := mux.Vars(r)["season"]
	season := Season{}

	err := h.DB.QueryRow(`select
		id, series_id, number, banner
		from seasons
		where id = (?)`, seasonID).Scan(&season.ID, &season.SeriesID, &season.Number, &season.Banner)

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
