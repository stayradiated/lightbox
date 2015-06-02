package db

import "time"

type Category struct {
	ID   int
	Name string
}

type Episode struct {
	ID                   int       `json:",omitempty"`
	SeasonID             int       `json:",omitempty"`
	ShowID               int       `json:",omitempty"`
	MediaID              int       `json:",omitempty"`
	DateCreated          time.Time `json:",omitempty"`
	DatePublished        time.Time `json:",omitempty"`
	Number               int       `json:",omitempty"`
	Title                string    `json:",omitempty"`
	Plot                 string    `json:",omitempty"`
	Runtime              int       `json:",omitempty"`
	FirstAired           time.Time `json:",omitempty"`
	Year                 int       `json:",omitempty"`
	ParentalRating       string    `json:",omitempty"`
	ParentalRatingReason string    `json:",omitempty"`
	Director             string    `json:",omitempty"`
	Writer               string    `json:",omitempty"`
	GuestStars           string    `json:",omitempty"`
	Rating               float64   `json:",omitempty"`
	RatingCount          int       `json:",omitempty"`
	Image                string    `json:",omitempty"`
	IMDB                 string    `json:",omitempty"`
	TVDB                 int       `json:",omitempty"`
}

type Season struct {
	ID             int       `json:",omitempty"`
	ShowID         int       `json:",omitempty"`
	DateCreated    time.Time `json:",omitempty"`
	Number         int       `json:",omitempty"`
	ParentalRating string    `json:",omitempty"`
	Image          string    `json:",omitempty"`
	TVDB           int       `json:",omitempty"`
	Episodes       []Episode `json:",omitempty"`
	EpisodeCount   int       `json:",omitempty"`
}

type Show struct {
	ID             int        `json:",omitempty"`
	Title          string     `json:",omitempty"`
	Year           int        `json:",omitempty"`
	Released       time.Time  `json:",omitempty"`
	Runtime        int        `json:",omitempty"`
	Writer         string     `json:",omitempty"`
	Actors         string     `json:",omitempty"`
	Plot           string     `json:",omitempty"`
	Poster         string     `json:",omitempty"`
	Fanart         string     `json:",omitempty"`
	Rating         float64    `json:",omitempty"`
	RatingCount    int        `json:",omitempty"`
	Categories     []Category `json:",omitempty"`
	DateCreated    time.Time  `json:",omitempty"`
	ParentalRating string     `json:",omitempty"`
	IMDB           string     `json:",omitempty"`
	TVDB           int        `json:",omitempty"`
	Seasons        []Season   `json:",omitempty"`
}
