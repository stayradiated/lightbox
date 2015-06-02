package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

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
			imdb_shows.title,
			lb_seasons.season_number,
			lb_episodes.episode_number
		from
			imdb_shows, lb_seasons, lb_episodes
		where
			imdb_shows.lightbox_id = lb_seasons.series_id and
			lb_seasons.id = lb_episodes.season_id
		order by
			lb_seasons.series_id
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var item struct {
		Title   string
		Season  string
		Episode string
	}

	for rows.Next() {
		err := rows.Scan(
			&item.Title,
			&item.Season,
			&item.Episode,
		)
		if err != nil {
			log.Fatal(err)
		}

		j, _ := json.Marshal(item)
		fmt.Println(string(j))
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
