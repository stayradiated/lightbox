package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var safety bool
	flag.BoolVar(&safety, "s", false, "safety run")
	flag.Parse()

	// connect to mysql
	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get all seasons that are not in the lightbox dataset
	rows, err := db.Query(`
		select seasons.id, seasons.number, shows.name
		from seasons, shows
		where
			shows.id = seasons.show_id and
			seasons.id not in (
				select seasons.id

				from seasons, shows, lb_seasons, lb_series

				where
					shows.lightbox_id = lb_series.id and
					shows.id = seasons.show_id and
					lb_series.id = lb_seasons.series_id and
					lb_seasons.season_number = seasons.number
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	deleteSeason, err := db.Prepare(`
		delete from seasons
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

		if safety == false {
			if _, err := deleteSeason.Exec(seasonID); err != nil {
				log.Fatal(err)
			}
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

}
