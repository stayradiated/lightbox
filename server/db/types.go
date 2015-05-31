package db

type Category string

type Episode struct {
	ID          int     `json:",omitempty"`
	Number      int     `json:",omitempty"`
	Name        string  `json:",omitempty"`
	Overview    string  `json:",omitempty"`
	FirstAired  string  `json:",omitempty"`
	Image       string  `json:",omitempty"`
	Rating      float64 `json:",omitempty"`
	RatingCount int     `json:",omitempty"`
	Director    string  `json:",omitempty"`
	Writer      string  `json:",omitempty"`
	GuestStars  string  `json:",omitempty"`
	IMDB        string  `json:",omitempty"`
}

type Season struct {
	ID       int       `json:",omitempty"`
	ShowID   int       `json:",omitempty"`
	Number   int       `json:",omitempty"`
	Episodes []Episode `json:",omitempty"`
	Banner   string    `json:",omitempty"`
}

type Show struct {
	ID            int        `json:",omitempty"`
	LightboxID    int        `json:",omitempty"`
	Name          string     `json:",omitempty"`
	Overview      string     `json:",omitempty"`
	Rating        float64    `json:",omitempty"`
	RatingCount   int        `json:",omitempty"`
	Categories    []Category `json:",omitempty"`
	Actors        string     `json:",omitempty"`
	Poster        string     `json:",omitempty"`
	Fanart        string     `json:",omitempty"`
	ContentRating string     `json:",omitempty"`
	FirstAired    string     `json:",omitempty"`
	Runtime       int        `json:",omitempty"`
	IMDB          string     `json:",omitempty"`
	Seasons       []Season   `json:",omitempty"`
}
