package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"bitbucket.org/stayradiated/lightbox/server/xstream"
)

type Image struct {
	ID     int
	Type   string
	Source string
}

func main() {

	var fp string

	flag.StringVar(&fp, "f", "", "path to json data")
	flag.Parse()

	if fp == "" {
		fmt.Println("WARNING: Must specify -f")
		return
	}

	db := DB{}
	db.Init()
	defer db.Close()

	f, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	seriesList := xstream.SeriesList{}

	j := json.NewDecoder(f)
	j.Decode(&seriesList)

	fmt.Printf("Loaded %d series...\n", len(seriesList))

	for _, series := range seriesList {
		insertImage(db.InsertShowImage, series.ID, series.Images)
		for _, season := range series.Seasons {
			insertImage(db.InsertSeasonImage, season.ID, season.Images)
			for _, episode := range season.Episodes {
				insertImage(db.InsertEpisodeImage, episode.ID, episode.Images)
				db.UpdateEpisodeDetails(episode.ID, episode.Details)
			}
		}
	}
}

type sqlFn func(int, Image)

func insertImage(fn sqlFn, id int, images []xstream.Image) {

	for _, image := range images {
		original, ok := image.Format["original_size"]

		if ok == true {
			image := Image{
				ID:     image.ID,
				Type:   image.Type,
				Source: original.Source,
			}

			fn(id, image)
		}
	}
}
