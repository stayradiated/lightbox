package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
)

func main() {

	// data, err := GetAllSeriesInCategory(COMEDY)
	// if err != nil {
	// 	panic(err)
	// }

	data, err := GetEpisodeInfo(356, 890, 12090)
	if err != nil {
		panic(err)
	}

	j, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(j))
}

func GetAllSeriesInCategory(categoryId int) (SeriesList, error) {
	series := make(SeriesList, 0)
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

func GetSeriesInCategory(categoryId, offset, limit int) (SeriesList, error) {
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

	response := CategoryResponse{}

	d := json.NewDecoder(r.Body)
	if err = d.Decode(&response); err != nil {
		return nil, err
	}

	return response.Series, nil
}

func GetSeriesInfo(seriesId int) (Series, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d",
		seriesId)

	response := Series{}

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

func GetSeasonInfo(seriesId, seasonId int) (Season, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d/seasons/%d",
		seriesId, seasonId)

	response := Season{}

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

func GetEpisodeInfo(seriesId, seasonId, episodeId int) (Episode, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d/seasons/%d/episodes/%d",
		seriesId, seasonId, episodeId)

	response := Episode{}

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
