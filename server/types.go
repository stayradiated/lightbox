package main

type Episode struct {
	ID          int
	Number      int
	Name        string
	Overview    string
	FirstAired  string
	Image       string
	Rating      float64
	RatingCount int
	Director    string
	Writer      string
	GuestStars  string
	IMDB        string
}

type Season struct {
	ID       int
	SeriesID int
	Number   int
	Episodes []Episode
}

type Series struct {
	ID            int      `json:",omitempty"`
	Name          string   `json:",omitempty"`
	Overview      string   `json:",omitempty"`
	Rating        float64  `json:",omitempty"`
	RatingCount   int      `json:",omitempty"`
	Genre         string   `json:",omitempty"`
	Actors        string   `json:",omitempty"`
	Poster        string   `json:",omitempty"`
	Banner        string   `json:",omitempty"`
	Fanart        string   `json:",omitempty"`
	ContentRating string   `json:",omitempty"`
	FirstAired    string   `json:",omitempty"`
	Runtime       int      `json:",omitempty"`
	IMDB          string   `json:",omitempty"`
	Seasons       []Season `json:",omitempty"`
}
