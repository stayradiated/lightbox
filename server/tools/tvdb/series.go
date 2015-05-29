package main

type SeriesInfo struct {
	Series   Series    `xml:"Series"`
	Episodes []Episode `xml:"Episode"`
}

type Series struct {
	ID            int    `xml:"seriesid"`
	Actors        string `xml:"Actors"`
	AirsDayOfWeek string `xml:"Airs_DayOfWeek"`
	AirsTime      string `xml:"Airs_Time"`
	ContentRating string `xml:"ContentRating"`
	FirstAired    string `xml:"FirstAired"`
	Genre         string `xml:"Genre"`
	IMDB          string `xml:"IMDB_ID"`
	Language      string `xml:"Language"`
	Network       string `xml:"Network"`
	NetworkID     string `xml:"NetworkID"`
	Overview      string `xml:"Overview"`
	Rating        string `xml:"Rating"`
	RatingCount   string `xml:"RatingCount"`
	Runtime       string `xml:"Runtime"`
	SeriesID      string `xml:"SeriesID"`
	SeriesName    string `xml:"SeriesName"`
	Status        string `xml:"Status"`
	Banner        string `xml:"banner"`
	FanArt        string `xml:"fanart"`
	LastUpdated   string `xml:"lastupdated"`
	Poster        string `xml:"poster"`
}

type Episode struct {
	ID                    int    `xml:"id"`
	CombinedEpisodeNumber string `xml:"Combined_episodenumber"`
	CombinedSeason        string `xml:"Combined_season"`
	Director              string `xml:"Director"`
	EpisodeName           string `xml:"EpisodeName"`
	EpisodeNumber         string `xml:"EpisodeNumber"`
	FirstAired            string `xml:"FirstAired"`
	GuestStars            string `xml:"GuestStars"`
	IMDB                  string `xml:"IMDB_ID"`
	Language              string `xml:"Language"`
	Overview              string `xml:"Overview"`
	Rating                string `xml:"Rating"`
	RatingCount           string `xml:"RatingCount"`
	SeasonNumber          string `xml:"SeasonNumber"`
	Writer                string `xml:"Writer"`
	AbsoluteNumber        string `xml:"absolute_number"`
	FileName              string `xml:"filename"`
	LastUpdated           string `xml:"lastupdated"`
	SeasonID              string `xml:"seasonid"`
	SeriesID              string `xml:"seriesid"`
}
