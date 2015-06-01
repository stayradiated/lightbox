package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"bitbucket.org/stayradiated/lightbox/server/xstream"
)

const (
	ALL_TV   = 9
	ALL_KIDS = 8
)

func main() {

	var everything bool
	var seriesID, seasonID, episodeID int

	flag.BoolVar(&everything, "everything", false, "everything")
	flag.IntVar(&seriesID, "series", -1, "series id")
	flag.IntVar(&seasonID, "season", -1, "season id")
	flag.IntVar(&episodeID, "episode", -1, "episode id")
	flag.Parse()

	var data interface{}
	var err error

	if everything && seriesID >= 0 {
		data, err = GetFullEpisodeInfo(seriesID)
	} else if everything == true {
		data, err = GetEverything()
	} else if episodeID >= 0 && seasonID >= 0 && seriesID >= 0 {
		data, err = GetEpisodeInfo(seriesID, seasonID, episodeID)
	} else if seasonID >= 0 && seriesID >= 0 {
		data, err = GetSeasonInfo(seriesID, seasonID)
	} else if seriesID >= 0 {
		data, err = GetSeriesInfo(seriesID)
	} else {
		fmt.Println("WARNING: Nothing to do...")
		return
	}

	if err != nil {
		panic(err)
	}

	j := json.NewEncoder(os.Stdout)
	j.Encode(data)
}

func GetEverything() (xstream.SeriesList, error) {
	seriesList, err := GetAllSeriesInCategory(ALL_TV)
	if err != nil {
		return nil, err
	}

	for i, series := range seriesList {
		seriesList[i], err = GetFullEpisodeInfo(series.ID)
	}

	return seriesList, nil
}

func GetFullEpisodeInfo(seriesID int) (series xstream.Series, err error) {

	series, err = GetSeriesInfo(seriesID)
	if err != nil {
		return series, err
	}

	for i, season := range series.Seasons {
		season, err = GetSeasonInfo(seriesID, season.ID)
		if err != nil {
			return series, err
		}
		series.Seasons[i] = season
	}

	return series, nil
}

func GetAllSeriesInCategory(categoryID int) (xstream.SeriesList, error) {
	series := make(xstream.SeriesList, 0)
	for {
		s, err := GetSeriesInCategory(categoryID, len(series), 50)
		if err != nil {
			return nil, err
		}
		if len(s) == 0 {
			break
		}
		series = append(series, s...)
	}
	return series, nil
}

func GetSeriesInCategory(categoryID, offset, limit int) (xstream.SeriesList, error) {
	baseurl := "https://www.lightbox.co.nz/xstream/media/series"

	v := url.Values{}
	v.Set("order", "asc")
	v.Set("sort", "title")
	v.Set("limit", strconv.Itoa(limit))
	v.Set("category_id", strconv.Itoa(categoryID))
	v.Set("offset", strconv.Itoa(offset))
	params := v.Encode()

	url := fmt.Sprintf("%s?%s", baseurl, params)

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	response := xstream.CategoryResponse{}

	d := json.NewDecoder(r.Body)
	if err = d.Decode(&response); err != nil {
		return nil, err
	}

	return response.Series, nil
}

func GetSeriesInfo(seriesID int) (xstream.Series, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d",
		seriesID)

	response := xstream.Series{}

	r, err := http.Get(url)
	if err != nil {
		return response, err
	}
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)
	if err = d.Decode(&response); err != nil {
		return response, err
	}

	return response, nil
}

func GetSeasonInfo(seriesID, seasonID int) (xstream.Season, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d/seasons/%d",
		seriesID, seasonID)

	response := xstream.Season{}

	r, err := http.Get(url)
	if err != nil {
		return response, err
	}
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)
	if err = d.Decode(&response); err != nil {
		return response, err
	}

	episodes, err := GetAllEpisodes(seriesID, seasonID)
	if err != nil {
		return response, err
	}
	response.Episodes = episodes

	return response, nil
}

type EpisodeList struct {
	Count    int
	Episodes []xstream.Episode
}

func GetAllEpisodes(seriesID, seasonID int) ([]xstream.Episode, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d/seasons/%d/episodes?limit=50",
		seriesID, seasonID)

	var response EpisodeList

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)
	if err = d.Decode(&response); err != nil {
		return nil, err
	}

	return response.Episodes, nil
}

func GetEpisodeInfo(seriesID, seasonID, episodeID int) (xstream.Episode, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d/seasons/%d/episodes/%d",
		seriesID, seasonID, episodeID)

	response := xstream.Episode{}

	r, err := http.Get(url)
	if err != nil {
		return response, err
	}
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)
	if err = d.Decode(&response); err != nil {
		return response, err
	}

	return response, nil
}
