package main

import "time"

type Show struct {
	ID          int
	Title       string
	Year        int
	Released    time.Time
	Runtime     int
	Writer      string
	Actors      string
	PLot        string
	Awards      string
	Poster      string
	Fanart      string
	Rating      float64
	RatingCount int
	DateCreated time.Time
	Categories  []int
	IMDB        string
	TVDB        int
}

func SmooshShows() {

	// loop through each show in lb_series

	// query related shows in imdb_shows and shows

	// make new Show object

	// set ID = lightbox id

	// set Title (IMDB, TVDB, Lightbox)

	// set Year (IMDB)

	// set Released (IMDB, TVDB)

	// set Runtime (IMDB, TVDB)

	// set Writer (IDMB)

	// set Actors (TVDB)

	// set Plot (IMDB, TVDB, Lightbox)

	// set Awards (IDMB)

	// set Poster (TVDB, IMDB, Lightbox)

	// set Fanart (TVDB, Lightbox?)

	// choose by highest rating count
	// else rating = 5 and count = 0
	// set Rating (IMDB, TVDB)
	// set RatingCount (IMDB, TVDB)

	// set DateCreated (Lightbox)

	// set Categories (IMDB, TVDB, Lightbox)

	// set IMDB (IMDB)

	// set TVDB (TVDB)

}
