package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"bitbucket.org/stayradiated/lightbox/server/db"
)

const baseUrl = "http://localhost:9000/"

func main() {

	var shows []db.Show
	if err := DownloadAndParseJSON("api/shows.json", &shows); err != nil {
		panic(err)
	}

	for _, show := range shows {

		var fullShow db.Show

		if err := DownloadAndParseJSON(
			fmt.Sprintf("api/shows/%d.json", show.ID),
			&fullShow,
		); err != nil {
			panic(err)
		}

		for _, season := range fullShow.Seasons {

			var fullSeason db.Season

			if err := DownloadAndParseJSON(
				fmt.Sprintf("api/seasons/%d.json", season.ID),
				&fullSeason,
			); err != nil {
				panic(err)
			}

			for _, episode := range fullSeason.Episodes {

				if err := DownloadFile(
					fmt.Sprintf("api/episodes/%d.json", episode.ID),
				); err != nil {
					panic(err)
				}

			}

		}

	}

	var categories []db.Category
	if err := DownloadAndParseJSON("api/categories.json", &categories); err != nil {
		panic(err)
	}

	for _, category := range categories {

		if err := DownloadFile(
			fmt.Sprintf("api/categories/%d.json", category.ID),
		); err != nil {
			panic(err)
		}

	}

}

func DownloadAndParseJSON(path string, data interface{}) error {
	r, err := http.Get(baseUrl + path)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	d := json.NewDecoder(io.TeeReader(r.Body, f))
	if err := d.Decode(data); err != nil {
		return err
	}

	return nil
}

func DownloadFile(path string) error {
	r, err := http.Get(baseUrl + path)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, r.Body); err != nil {
		return err
	}

	return nil
}
