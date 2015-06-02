package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DB *sql.DB

	insertShow    *sql.Stmt
	insertEpisode *sql.Stmt
}

func (d *DB) Init() {
	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}

	if d.insertShow, err = db.Prepare(`
		insert ignore into imdb_shows (
			lightbox_id,
			title, year, rated, released, runtime, genre, director, writer, 
			actors, plot, language, country, awards, poster, metascore, 
			imdb_rating, imdb_votes, imdb_id, dvd, boxoffice, production, website 
		)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )
	`); err != nil {
		log.Fatal(err)
	}

	if d.insertEpisode, err = db.Prepare(`
		insert ignore into imdb_episodes (
			title, year, rated, released, season, episode, runtime, genre, director,
			writer, actors, plot, language, country, awards, poster, metascore,
			imdb_rating, imdb_votes, imdb_id, show_id
		)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )
	`); err != nil {
		log.Fatal(err)
	}

	d.DB = db
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) InsertShow(lightbox int, show Show) {

	if _, err := d.insertShow.Exec(
		lightbox,
		show.Title,
		show.Year,
		show.Rated,
		show.Released,
		show.Runtime,
		show.Genre,
		show.Director,
		show.Writer,
		show.Axctors,
		show.Plot,
		show.Language,
		show.Country,
		show.Awards,
		show.Poster,
		show.Metascore,
		show.ImdbRating,
		show.ImdbVotes,
		show.ImdbID,
		show.DVD,
		show.BoxOffice,
		show.Production,
		show.Website,
	); err != nil {
		log.Fatal(err)
	}

}

func (d *DB) InsertEpisode(episode Episode) {

	if _, err := d.insertEpisode.Exec(
		episode.Title,
		episode.Year,
		episode.Rated,
		episode.Released,
		episode.Season,
		episode.Episode,
		episode.Runtime,
		episode.Genre,
		episode.Director,
		episode.Writer,
		episode.Actors,
		episode.Plot,
		episode.Language,
		episode.Country,
		episode.Awards,
		episode.Poster,
		episode.Metascore,
		episode.ImdbRating,
		episode.ImdbVotes,
		episode.ImdbID,
		episode.SeriesID,
	); err != nil {
		log.Fatal(err)
	}

}
