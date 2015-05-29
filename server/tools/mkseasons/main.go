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

	// prepare query
	insertSeason, err := db.Prepare(`INSERT IGNORE INTO seasons_tvdb(
		id,
		series_id,
		number
	) VALUES(?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	// Get All Episodes
	rows, err := db.Query("select series_id, season_id, season_number from episodes_tvdb")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		seriesID     int
		seasonID     int
		seasonNumber int
	)

	for rows.Next() {
		err := rows.Scan(&seriesID, &seasonID, &seasonNumber)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := insertSeason.Exec(
			seasonID,
			seriesID,
			seasonNumber,
		); err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
