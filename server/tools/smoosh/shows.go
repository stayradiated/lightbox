package main

import (
	"strconv"
	"time"
)

type Show struct {
	ID             int       `json:",omitempty"`
	Title          string    `json:",omitempty"`
	Year           int       `json:",omitempty"`
	Released       time.Time `json:",omitempty"`
	Runtime        int       `json:",omitempty"`
	Writer         string    `json:",omitempty"`
	Actors         string    `json:",omitempty"`
	Plot           string    `json:",omitempty"`
	Poster         string    `json:",omitempty"`
	Fanart         string    `json:",omitempty"`
	Rating         float64   `json:",omitempty"`
	RatingCount    int       `json:",omitempty"`
	DateCreated    time.Time `json:",omitempty"`
	ParentalRating string    `json:",omitempty"`
	IMDB           string    `json:",omitempty"`
	TVDB           int       `json:",omitempty"`
}

type ShowData struct {
	ID   int
	IMDB struct {
		ID       []byte
		Title    []byte
		Year     []byte
		Rated    []byte
		Released []byte
		Runtime  []byte
		Genre    []byte
		Writer   []byte
		Plot     []byte
		Language []byte
		Country  []byte
		Poster   []byte
		Rating   []byte
		Votes    []byte
		IMDB     []byte
	}
	TVDB struct {
		ID            []byte
		Name          []byte
		Overview      []byte
		Genre         []byte
		Actors        []byte
		ContentRating []byte
		FirstAired    []byte
		Language      []byte
		Network       []byte
		Rating        []byte
		RatingCount   []byte
		Runtime       []byte
		Status        []byte
		Banner        []byte
		Fanart        []byte
		Poster        []byte
	}
	LB struct {
		DateCreated    []byte
		Title          []byte
		Description    []byte
		ParentalRating []byte
		Image          []byte
	}
}

func (d *DB) SmooshShows() {

	for _, s := range d.GetAllShows() {

		imdb := s.IMDB
		tvdb := s.TVDB
		lightbox := s.LB

		// make new Show object
		show := Show{
			ID: s.ID,
		}

		// set Title (IMDB, TVDB, Lightbox)
		if isset(imdb.Title) {
			show.Title = string(imdb.Title)
		} else if isset(tvdb.Name) {
			show.Title = string(tvdb.Name)
		} else {
			show.Title = string(lightbox.Title)
		}

		// set Year (IMDB)
		if isset(imdb.Year) {
			show.Year, _ = strconv.Atoi(string(imdb.Year[0:4]))
		} else if isset(tvdb.FirstAired) {
			show.Year, _ = strconv.Atoi(string(tvdb.FirstAired[0:4]))
		}

		// set Released (IMDB, TVDB)
		if isset(imdb.Released) {
			show.Released, _ = time.Parse("02 Jan 2006", string(imdb.Released))
		} else if isset(tvdb.FirstAired) {
			show.Released = parseDate(tvdb.FirstAired)
		}

		// set Runtime (IMDB, TVDB)
		if isset(imdb.Runtime) {
			show.Runtime = atoi(imdb.Runtime)
		} else if isset(tvdb.Runtime) {
			show.Runtime = atoi(tvdb.Runtime)
		}

		// set Writer (IDMB)
		if isset(imdb.Writer) {
			show.Writer = string(imdb.Writer)
		}

		// set Actors (TVDB)
		if isset(tvdb.Actors) {
			show.Actors = unpipe(tvdb.Actors)
		}

		// set Plot (IMDB, TVDB, Lightbox)
		if isset(imdb.Plot) {
			show.Plot = string(imdb.Plot)
		} else if isset(tvdb.Overview) {
			show.Plot = string(tvdb.Overview)
		} else {
			show.Plot = string(lightbox.Description)
		}

		// set Poster (TVDB, IMDB, Lightbox)
		if isset(tvdb.Poster) {
			show.Poster = tvdbBannerURL + string(tvdb.Poster)
		} else if isset(imdb.Poster) {
			show.Poster = string(imdb.Poster)
		} else {
			show.Poster = string(lightbox.Image)
		}

		// set Fanart (TVDB, Lightbox?)
		if isset(tvdb.Fanart) {
			show.Fanart = tvdbBannerURL + string(tvdb.Fanart)
		}

		// set Rating, RatingCount (IMDB, TVDB)
		if isset(imdb.Votes) {
			show.Rating, _ = strconv.ParseFloat(string(imdb.Rating), 64)
			show.RatingCount = uncomma(imdb.Votes)
		} else if isset(tvdb.RatingCount) {
			show.Rating, _ = strconv.ParseFloat(string(tvdb.Rating), 64)
			show.RatingCount, _ = strconv.Atoi(string(tvdb.RatingCount))
		} else {
			show.Rating = 5
			show.RatingCount = 0
		}

		// set DateCreated (Lightbox)
		if isset(lightbox.DateCreated) {
			show.DateCreated = parseDateTime(lightbox.DateCreated)
		}

		// set ParentalRating (Lightbox)
		if isset(lightbox.ParentalRating) {
			show.ParentalRating = string(lightbox.ParentalRating)
		}

		// set IMDB (IMDB)
		if isset(imdb.IMDB) {
			show.IMDB = string(imdb.IMDB)
		}

		// set TVDB (TVDB)
		if isset(tvdb.ID) {
			show.TVDB, _ = strconv.Atoi(string(tvdb.ID))
		}

		d.InsertMasterShow(show)
	}

}
