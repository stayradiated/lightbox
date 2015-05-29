package main

import (
	"bitbucket.org/stayradiated/lightbox/server/xstream"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	COMEDY      = 18
	CRIME       = 22
	DRAMA       = 7
	FACTUAL     = 17
	NEW_ZEALAND = 1
	REALITY     = 31
	SCIFI       = 16
	ALL_TV      = 9
	PRE_SCHOOL  = 33
	ALL_KIDS    = 8
)

func main() {

	var everything bool
	var seriesId, seasonId, episodeId int

	flag.BoolVar(&everything, "everything", false, "get everything")
	flag.IntVar(&seriesId, "series", -1, "series id")
	flag.IntVar(&seasonId, "season", -1, "season id")
	flag.IntVar(&episodeId, "episode", -1, "episode id")
	flag.Parse()

	var data interface{}
	var err error

	if everything == true {
		data, err = GetEverything()
	} else if episodeId >= 0 && seasonId >= 0 && seriesId >= 0 {
		data, err = GetEpisodeInfo(seriesId, seasonId, episodeId)
	} else if seasonId >= 0 && seriesId >= 0 {
		data, err = GetSeasonInfo(seriesId, seasonId)
	} else if seriesId >= 0 {
		data, err = GetSeriesInfo(seriesId)
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
		fmt.Println(series.Id, series.Titles.Default)

		series, err = GetSeriesInfo(series.Id)
		if err != nil {
			return nil, err
		}

		seriesList[i] = series

		for j, season := range series.Seasons {
			fmt.Println("  -", season.Titles.Default)

			season, err = GetSeasonInfo(series.Id, season.Id)
			if err != nil {
				return nil, err
			}

			series.Seasons[j] = season
		}
	}

	return seriesList, nil
}

func GetAllSeriesInCategory(categoryId int) (xstream.SeriesList, error) {
	series := make(xstream.SeriesList, 0)
	for {
		s, err := GetSeriesInCategory(categoryId, len(series), 50)
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

func GetSeriesInCategory(categoryId, offset, limit int) (xstream.SeriesList, error) {
	baseurl := "https://www.lightbox.co.nz/xstream/media/series"

	v := url.Values{}
	v.Set("order", "asc")
	v.Set("sort", "title")
	v.Set("limit", strconv.Itoa(limit))
	v.Set("category_id", strconv.Itoa(categoryId))
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

func GetSeriesInfo(seriesId int) (xstream.Series, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d",
		seriesId)

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

func GetSeasonInfo(seriesId, seasonId int) (xstream.Season, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d/seasons/%d",
		seriesId, seasonId)

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

	return response, nil
}

func GetEpisodeInfo(seriesId, seasonId, episodeId int) (xstream.Episode, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d/seasons/%d/episodes/%d",
		seriesId, seasonId, episodeId)

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
