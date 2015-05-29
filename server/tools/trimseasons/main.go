package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// connect to mysql
	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get all seasons that are not in the lightbox dataset
	rows, err := db.Query(`
		select seasons_tvdb.id, seasons_tvdb.number, series_tvdb.series_name
		from seasons_tvdb, series_tvdb
		where

		series_tvdb.id = seasons_tvdb.series_id
		and seasons_tvdb.id not in (
		select seasons_tvdb.id

		from seasons_tvdb, seasons, series, series_tvdb

		where
			series_tvdb.lightbox_id = series.id and
			series_tvdb.id = seasons_tvdb.series_id and
			series.id = seasons.series_id and
			seasons.season_number = seasons_tvdb.number
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	deleteSeason, err := db.Prepare(`
		delete from seasons_tvdb
		where id = ?`)
	if err != nil {
		log.Fatal(err)
	}

	var (
		seasonID     int
		seasonNumber int
		seriesTitle  string
	)

	for rows.Next() {
		if err := rows.Scan(&seasonID, &seasonNumber, &seriesTitle); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Deleting", seriesTitle, seasonNumber)
		if _, err := deleteSeason.Exec(seasonID); err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

}
