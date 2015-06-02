package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	database "bitbucket.org/stayradiated/lightbox/server/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox_backup")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		select
			id, lightbox_id, name, overview, rating, rating_count, actors, poster,
			fanart, content_rating, first_aired, runtime, imdb
		from shows
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var show database.Show

	for rows.Next() {
		err := rows.Scan(
			&show.ID,
			&show.LightboxID,
			&show.Name,
			&show.Overview,
			&show.Rating,
			&show.RatingCount,
			&show.Actors,
			&show.Poster,
			&show.Fanart,
			&show.ContentRating,
			&show.FirstAired,
			&show.Runtime,
			&show.IMDB,
		)
		if err != nil {
			log.Fatal(err)
		}

		j, _ := json.Marshal(show)
		fmt.Println(string(j))
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
