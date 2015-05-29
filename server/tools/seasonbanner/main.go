package main

import (
	"database/sql"
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

	// prepare update query
	updateSeasonBanner, err := db.Prepare(`
		update seasons_tvdb
		set banner = ?
		where id = ?`)
	if err != nil {
		log.Fatal(err)
	}

	// get all seasons
	rows, err := db.Query(`
		select sr.id, sn.number, sn.id, b.banner_path

		from
			series_tvdb sr,
				seasons_tvdb sn,
				banners b

		where
			sr.id = sn.series_id and
				b.series_id = sr.id and
				b.season = sn.number and
				b.banner_type = 2 and
				b.banner_size = "season" and
				b.language = "en"
				
		group by sr.id, sn.number
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		seriesID     int
		seasonNumber int
		seasonID     int
		banner       string
	)

	for rows.Next() {

		err := rows.Scan(&seriesID, &seasonNumber, &seasonID, &banner)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := updateSeasonBanner.Exec(banner, seasonID); err != nil {
			log.Fatal(err)
		}

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
