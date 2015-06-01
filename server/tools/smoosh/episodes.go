package main

import (
	"strconv"
	"time"
)

type Episode struct {
	ID                   int
	SeasonID             int
	MediaID              int
	DateCreated          time.Time
	DatePublished        time.Time
	Number               int
	Title                string
	Plot                 string
	Runtime              int
	FirstAired           time.Time
	Year                 int
	ParentalRating       string
	ParentalRatingReason string
	Director             string
	Writer               string
	GuestStars           string
	Rating               float64
	RatingCount          int
	Image                string
	IMDB                 string
	TVDB                 int
}

type EpisodeData struct {
	ID   int
	TVDB struct {
		ID           []byte
		ShowID       []byte
		SeasonID     []byte
		Director     []byte
		Name         []byte
		Number       []byte
		FirstAired   []byte
		GuestStars   []byte
		IMDB         []byte
		Language     []byte
		Overview     []byte
		Rating       []byte
		RatingCount  []byte
		SeasonNumber []byte
		Writer       []byte
		Filename     []byte
	}
	LB struct {
		SeasonID       []byte
		DatePublished  []byte
		DateCreated    []byte
		Title          []byte
		Description    []byte
		ParentalRating []byte
		Number         []byte
		MediaID        []byte
		Runtime        []byte
		AirDate        []byte
		NZRating       []byte
		NZRatingReason []byte
		SeriesCast     []byte
		Year           []byte
	}
}

func (d *DB) SmooshEpisodes() {

	for _, e := range d.GetAllEpisodes() {

		tvdb := e.TVDB
		lightbox := e.LB

		episode := Episode{
			ID: e.ID,
		}

		// Set SeasonID
		if isset(lightbox.SeasonID) {
			episode.SeasonID = atoi(lightbox.SeasonID)
		}

		// Set MediaID
		if isset(lightbox.MediaID) {
			episode.MediaID = atoi(lightbox.MediaID)
		}

		// Set DateCreated
		if isset(lightbox.DateCreated) {
			episode.DateCreated = parseDateTime(lightbox.DateCreated)
		}

		// Set DatePublished
		if isset(lightbox.DatePublished) {
			episode.DatePublished = parseDateTime(lightbox.DatePublished)
		}

		// Set Number
		if isset(lightbox.Number) {
			episode.Number = atoi(lightbox.Number)
		} else if isset(tvdb.Number) {
			episode.Number = atoi(tvdb.Number)
		}

		// Set Title
		if isset(tvdb.Name) {
			episode.Title = string(tvdb.Name)
		} else if isset(lightbox.Title) {
			episode.Title = string(lightbox.Title)
		}

		// Set Plot
		if isset(tvdb.Overview) {
			episode.Plot = string(tvdb.Overview)
		} else if isset(lightbox.Description) {
			episode.Plot = string(lightbox.Description)
		}

		// Set Runtime
		if isset(lightbox.Runtime) {
			episode.Runtime = atoi(lightbox.Runtime)
		}

		// Set FirstAired
		if isset(lightbox.AirDate) {
			episode.FirstAired = parseDateTime(lightbox.AirDate)
		} else if isset(tvdb.FirstAired) {
			episode.FirstAired = parseDate(tvdb.FirstAired)
		}

		// Set Year
		episode.Year = episode.FirstAired.Year()

		// Set ParentalRating
		if isset(lightbox.NZRating) {
			episode.ParentalRating = string(lightbox.NZRating)
		} else if isset(lightbox.ParentalRating) {
			episode.ParentalRating = string(lightbox.ParentalRating)
		}

		// Set ParentalRatingReason
		if isset(lightbox.NZRatingReason) {
			episode.ParentalRatingReason = string(lightbox.NZRatingReason)
		}

		// Set Director
		if isset(tvdb.Director) {
			episode.Director = unpipe(tvdb.Director)
		}

		// Set Writer
		if isset(tvdb.Writer) {
			episode.Writer = unpipe(tvdb.Writer)
		}

		// Set GuestStars
		if isset(tvdb.GuestStars) {
			episode.GuestStars = unpipe(tvdb.GuestStars)
		}

		// Set Rating
		if isset(tvdb.Rating) {
			episode.Rating, _ = strconv.ParseFloat(string(tvdb.Rating), 64)
		} else {
			episode.Rating = 5
		}

		// Set RatingCount
		if isset(tvdb.RatingCount) {
			episode.RatingCount = atoi(tvdb.RatingCount)
		} else {
			episode.RatingCount = 0
		}

		// Set Image
		if isset(tvdb.Filename) {
			episode.Image = tvdbBannerURL + string(tvdb.Filename)
		}

		// Set IMDB
		if isset(tvdb.IMDB) {
			episode.IMDB = string(tvdb.IMDB)
		}

		// Set TVDB
		if isset(tvdb.ID) {
			episode.TVDB = atoi(tvdb.ID)
		}

		d.InserMasterEpisode(episode)
	}

}
