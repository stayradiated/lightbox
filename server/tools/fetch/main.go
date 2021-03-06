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

/*

1: Most Popular (Dynamic)
2: Comedy Picks (Dynamic)
3: Lightbox Kids (Dynamic)
4:
5: Most Popular TV Drama (Dynamic)
6: Most Popular TV Comedy (Dynamic)
7: Most Popular TV Crime (Dynamic)
8: Most Popular TV Sci Fi/Fantasy (dynamic)
9: Most Popular TV Reality
10: Most Popular Factual
11: Most Popular Pre-school
12: Most Popular All kids
13:
14: Emmy Nomiated
15: Sci-Fi Wonders
16: Staff Obsessions
17: Good For A Laugh
18: School Holidays
19: Recently Added
20: Most Popular New Zealand
21: Lightbox Exclusives
22: British Crime
23: British Comedy
24: Staff Picks

*/

func main() {

	var everything, hotspots bool
	var seriesID, seasonID, episodeID, listID int

	flag.BoolVar(&everything, "everything", false, "everything")
	flag.BoolVar(&hotspots, "hotspots", false, "hotspots")
	flag.IntVar(&seriesID, "series", -1, "series id")
	flag.IntVar(&seasonID, "season", -1, "season id")
	flag.IntVar(&episodeID, "episode", -1, "episode id")
	flag.IntVar(&listID, "list", -1, "list id")
	flag.Parse()

	var data interface{}
	var err error

	if listID >= 0 {
		data, err = GetSectionList(listID)
	} else if hotspots {
		// TODO: implement hotspots
	} else if everything && seriesID >= 0 {
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

func GetEverything() ([]xstream.Series, error) {
	seriesList, err := GetAllSeriesInCategory(ALL_KIDS)
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

func GetAllSeriesInCategory(categoryID int) ([]xstream.Series, error) {
	series := make([]xstream.Series, 0)
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

func GetSeriesInCategory(categoryID, offset, limit int) ([]xstream.Series, error) {
	baseurl := "https://www.lightbox.co.nz/xstream/media/series"

	v := url.Values{}
	v.Set("order", "asc")
	v.Set("sort", "title")
	v.Set("limit", strconv.Itoa(limit))
	v.Set("category_id", strconv.Itoa(categoryID))
	v.Set("offset", strconv.Itoa(offset))
	params := v.Encode()

	url := fmt.Sprintf("%s?%s", baseurl, params)

	response := xstream.SeriesList{}
	if err := request(url, &response); err != nil {
		return response.Series, err
	}

	return response.Series, nil
}

func GetSeriesInfo(seriesID int) (xstream.Series, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d",
		seriesID)

	response := xstream.Series{}
	if err := request(url, &response); err != nil {
		return response, err
	}

	return response, nil
}

func GetSeasonInfo(seriesID, seasonID int) (xstream.Season, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d/seasons/%d",
		seriesID, seasonID)

	response := xstream.Season{}
	if err := request(url, &response); err != nil {
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
	if err := request(url, &response); err != nil {
		return nil, err
	}

	return response.Episodes, nil
}

func GetEpisodeInfo(seriesID, seasonID, episodeID int) (xstream.Episode, error) {
	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/media/series/%d/seasons/%d/episodes/%d",
		seriesID, seasonID, episodeID)

	response := xstream.Episode{}
	if err := request(url, &response); err != nil {
		return response, err
	}

	return response, nil
}

func GetSectionList(listID int) (xstream.List, error) {

	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/sections/lists/%d",
		listID,
	)

	response := xstream.List{}
	if err := request(url, &response); err != nil {
		return response, err
	}

	elements, err := GetSectionListElements(listID)
	if err != nil {
		return response, err
	}
	response.Elements = elements

	return response, nil
}

func GetSectionListElements(listID int) (xstream.SeriesList, error) {

	url := fmt.Sprintf(
		"https://www.lightbox.co.nz/xstream/sections/lists/%d/elements?limit=50",
		listID,
	)

	var response xstream.ListElements
	if err := request(url, &response); err != nil {
		return response.Elements, err
	}

	return response.Elements, nil
}

func request(url string, data interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)
	if err = d.Decode(data); err != nil {
		return err
	}

	return nil
}
