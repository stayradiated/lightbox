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
	if limit > 50 {
		limit = 50
	}

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if offset == 0 || err != nil {
		offset = 0
	}

	rows, err := h.DB.Query(`select
		id, series_name, overview, poster
		from series_tvdb
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
		err := rows.Scan(&series.ID, &series.Name, &series.Overview, &series.Poster)
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
		genre,
		actors,
		poster,
		banner,
		fanart,
		content_rating,
		first_aired,
		runtime,
		imdb
	from series_tvdb
	where id = (?)`, seriesID).Scan(
		&series.ID,
		&series.Name,
		&series.Overview,
		&series.Rating,
		&series.RatingCount,
		&series.Genre,
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

	rows, err := h.DB.Query(`select
		id, series_id, number
		from seasons_tvdb
		where series_id = (?)`, seriesID)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	seasons := make([]Season, 0)

	for rows.Next() {
		season := Season{}
		err = rows.Scan(&season.ID, &season.SeriesID, &season.Number)
		if err != nil {
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
		id, series_id, number
		from seasons_tvdb
		where id = (?)`, seasonID).Scan(&season.ID, &season.SeriesID, &season.Number)

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
	from episodes_tvdb where season_id = (?)`, seasonID)

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
