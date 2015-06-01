package main

import "time"

type Season struct {
	ID             int
	ShowID         int
	DateCreated    time.Time
	Number         int
	ParentalRating string
	Image          string
	TVDB           int
}

type SeasonData struct {
	ID   int
	TVDB struct {
		ID     []byte
		ShowID []byte
		Number []byte
		Banner []byte
	}
	LB struct {
		ShowID         []byte
		DateCreated    []byte
		Number         []byte
		Title          []byte
		Description    []byte
		ParentalRating []byte
		Image          []byte
	}
}

func (d *DB) SmooshSeasons() {

	for _, s := range d.GetAllSeasons() {

		tvdb := s.TVDB
		lightbox := s.LB

		season := Season{
			ID: s.ID,
		}

		// set ShowID (Lightbox)
		if isset(lightbox.ShowID) {
			season.ShowID = atoi(lightbox.ShowID)
		}

		// set DateCreated (Lightbox)
		if isset(lightbox.DateCreated) {
			season.DateCreated = parseDateTime(lightbox.DateCreated)
		}

		// set Number (Lightbox)
		if isset(lightbox.Number) {
			season.Number = atoi(lightbox.Number)
		}

		// set ParentalRating (Lightbox)
		if isset(lightbox.ParentalRating) {
			season.ParentalRating = string(lightbox.ParentalRating)
		}

		// set Image (TVDB, Lightbox)
		if isset(tvdb.Banner) {
			season.Image = tvdbBannerURL + string(tvdb.Banner)
		} else if isset(lightbox.Image) {
			season.Image = string(lightbox.Image)
		}

		// set TVDB (TVDB)
		if isset(tvdb.ID) {
			season.TVDB = atoi(tvdb.ID)
		}

		d.InsertMasterSeason(season)
	}

}
