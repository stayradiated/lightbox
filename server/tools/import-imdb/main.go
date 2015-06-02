package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	var id, show, season, episode string
	var lightbox int

	flag.IntVar(&lightbox, "lightbox", 0, "lightbox id")
	flag.StringVar(&id, "id", "", "IMDB ID")
	flag.StringVar(&show, "show", "", "Show Name")
	flag.StringVar(&season, "season", "", "Season Number")
	flag.StringVar(&episode, "episode", "", "Episode Number")
	flag.Parse()

	db := DB{}
	db.Init()
	defer db.Close()

	if len(show) > 0 && len(episode) > 0 && len(season) > 0 {
		data, err := getEpisode(show, season, episode)
		if err != nil {
			log.Fatal(err)
		}
		j, _ := json.Marshal(data)
		fmt.Println(string(j))
		// db.InsertEpisode(data)
		return
	}

	if len(id) > 0 {
		fmt.Println("Getting show", id, lightbox)
		data, err := getShow(id)
		if err != nil {
			log.Fatal(err)
		}
		db.InsertShow(lightbox, data)
		return
	}

}

func getEpisode(show, season, episode string) (Episode, error) {

	v := url.Values{}
	v.Set("t", show)
	v.Set("Season", season)
	v.Set("Episode", episode)

	fmt.Println(v.Encode())

	data := Episode{}

	if err := callOMDb(&data, v.Encode()); err != nil {
		return data, err
	}

	return data, nil
}

func getShow(id string) (Show, error) {

	v := url.Values{}
	v.Set("i", id)
	v.Set("plot", "full")

	show := Show{}

	if err := callOMDb(&show, v.Encode()); err != nil {
		return show, err
	}

	return show, nil
}

func callOMDb(data interface{}, params string) error {

	url := fmt.Sprintf("http://www.omdbapi.com/?%s", params)

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)
	if err := d.Decode(data); err != nil {
		return err
	}

	return nil
}
