package xstream

type ImageFormat struct {
	Format string `json:"format"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Source string `json:"source"`
}

type Image struct {
	ID     int                    `json:"id"`
	Format map[string]ImageFormat `json:"format"`
	Type   string                 `json:"type"`
}

type MapString struct {
	Default string `json:"default"`
	En      string `json:"en"`
}

type Dates struct {
	Published string `json:"published"`
	Created   string `json:"created"`
}

type Category struct {
	ID     int       `json:"id"`
	Parent int       `json:"parent"`
	Titles MapString `json:"titles"`
	// Children
}

type ParentalControl struct {
	Rating string `json:"rating"`
}

type Details struct {
	Length                   string   `json:"length"`
	AirDate                  string   `json:"air_date"`
	ContentOwner             string   `json:"content_owner"`
	TmsVersionID             string   `json:"tms_version_id"`
	TmsAltFilmID             string   `json:"tms_alt_film_id"`
	TmsRootID                string   `json:"tms_root_id"`
	OrderID                  string   `json:"order_id"`
	SeriesTmsID              string   `json:"series_tms_id"`
	SeriesTmsSeriesID        string   `json:"series_tms_series_id"`
	SeriesTmsRootID          string   `json:"series_tms_root_id"`
	SeriesNzRating           string   `json:"series_nz_rating"`
	SeriesNzRatingAdvisories string   `json:"series_nz_rating_advisories"`
	SeriesNzRatingReason     string   `json:"series_nz_rating_reason"`
	SeriesAirDate            string   `json:"series_air_date"`
	SeriesCast               []string `json:"series_cast"`
	// SeriesAward
	Year               string   `json:"year"`
	RegID              string   `json:"reg_id"`
	Director           []string `json:"director"`
	NzRatingAdvisories string   `json:"nz_rating_advisories"`
	NzRatingReason     string   `json:"nz_rating_reason"`
	NzRating           string   `json:"nz_rating"`
}

type Stream struct {
	ID      int    `json:"id"`
	Bitrate int    `json:"bitrate"`
	Src     string `json:"src"`
	Type    string `json:"type"`
	Size    int    `json:"size"`
	Drm     struct {
		Type string `json:"type"`
	} `json:"drm"`
	Flags []string `json:"flags"`
}

type Episode struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	Titles           MapString `json:"titles"`
	Dates            Dates     `json:"dates"`
	Descriptions     MapString `json:"descriptions"`
	LongDescriptions MapString `json:"long_descriptions"`
	Images           []Image   `json:"images"`
	// Stats
	// Series
	// Pricing
	Details         Details         `json:"details"`
	ParentalControl ParentalControl `json:"parental_control"`
	ContentProvider struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Title string `json:"title"`
	} `json:"content_provider"`
	Categories []Category `json:"categories"`
	Episode    int        `json:"episode"`
	Streams    struct {
		Web []Stream `json:"web"`
	} `json:"streams"`
	Purchased bool `json:"purchased"`
	MediaID   int  `json:"media_id"`
	SeriesID  int  `json:"series_id"`
	SeasonID  int  `json:"season_id"`
}

type Season struct {
	ID           int       `json:"id"`
	DateCreated  string    `json:"date_created"`
	SeasonNumber int       `json:"season_number"`
	Episodes     []Episode `json:"episodes"`
	EpisodeCount int       `json:"episode_count"`
	Images       []Image   `json:"images"`
	Titles       MapString `json:"titles"`
	// Trailers
	Categories       []Category      `json:"categories"`
	Descriptions     MapString       `json:"descriptions"`
	LongDescriptions MapString       `json:"long_descriptions"`
	ParentalControl  ParentalControl `json:"parental_control"`
}

type Series struct {
	ID               int             `json:"id"`
	DateCreated      string          `json:"date_created"`
	Images           []Image         `json:"images"`
	Titles           MapString       `json:"titles"`
	Descriptions     MapString       `json:"descriptions"`
	LongDescriptions MapString       `json:"long_descriptions"`
	Seasons          []Season        `json:"seasons"`
	Categories       []Category      `json:"categories"`
	Type             string          `json:"type"`
	ParentalControl  ParentalControl `json:"parental_control"`
}

type SeriesList struct {
	Series []Series `json:"series"`
	Count  int      `json:"count"`
}
