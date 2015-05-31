package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"bitbucket.org/stayradiated/lightbox/server/xstream"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var fp string

	flag.StringVar(&fp, "f", "", "path to json data")
	flag.Parse()

	if fp == "" {
		fmt.Println("WARNING: Must specify -f")
		return
	}

	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	insertSeries, err := db.Prepare(`INSERT IGNORE INTO lb_series(
		id,
		date_created,
		title,
		description,
		parental_rating
	) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	insertSeriesCategory, err := db.Prepare(`INSERT IGNORE INTO lb_series_categories(
		series_id,
		category_id
	) VALUES(?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	insertSeason, err := db.Prepare(`INSERT IGNORE INTO lb_seasons(
		id,
		series_id,
		date_created,
		season_number,
		title,
		description,
		parental_rating
	) VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	insertEpisode, err := db.Prepare(`INSERT IGNORE INTO lb_episodes(
		id,
		season_id,
		date_published,
		date_created,
		title,
		description,
		parental_rating,
		episode_number,
		media_id,
		runtime,
		air_date
	) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

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

		_, err = insertSeries.Exec(
			series.ID,
			series.DateCreated,
			series.Titles.Default,
			series.Descriptions.Default,
			series.ParentalControl.Rating,
		)

		if err != nil {
			log.Fatal(err)
		}

		for _, category := range series.Categories {

			_, err := insertSeriesCategory.Exec(
				series.ID,
				category.ID,
			)

			if err != nil {
				log.Fatal(err)
			}

		}

		for _, season := range series.Seasons {

			_, err := insertSeason.Exec(
				season.ID,
				series.ID,
				season.DateCreated,
				season.SeasonNumber,
				season.Titles.Default,
				season.Descriptions.Default,
				season.ParentalControl.Rating,
			)

			if err != nil {
				log.Fatal(err)
			}

			for _, episode := range season.Episodes {

				_, err := insertEpisode.Exec(
					episode.ID,
					season.ID,
					episode.Dates.Published,
					episode.Dates.Created,
					episode.Titles.Default,
					episode.Descriptions.Default,
					episode.ParentalControl.Rating,
					episode.Episode,
					episode.MediaID,
					episode.Details.Length,
					episode.Details.AirDate,
				)

				if err != nil {
					log.Fatal(err)
				}

			}

		}

	}

}
