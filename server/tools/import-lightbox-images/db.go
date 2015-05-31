package main

import (
	"database/sql"
	"log"
	"strings"

	"bitbucket.org/stayradiated/lightbox/server/xstream"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DB *sql.DB

	insertShowImage    *sql.Stmt
	insertSeasonImage  *sql.Stmt
	insertEpisodeImage *sql.Stmt

	updateEpisodeDetails *sql.Stmt
}

func (d *DB) Init() {
	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}

	if d.insertShowImage, err = db.Prepare(`
		insert ignore into lb_show_images (
			id, show_id, type, source
		)
		values (?, ?, ?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	if d.insertSeasonImage, err = db.Prepare(`
		insert ignore into lb_season_images (
			id, season_id, type, source
		)
		values (?, ?, ?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	if d.insertEpisodeImage, err = db.Prepare(`
		insert ignore into lb_episode_images (
			id, episode_id, type, source
		)
		values (?, ?, ?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	if d.updateEpisodeDetails, err = db.Prepare(`
		update lb_episodes
		set
			year = ?,
			series_cast = ?,
			nz_rating = ?,
			nz_rating_reason = ?,
			nz_rating_advisories = ?
		where id = ?
	`); err != nil {
		log.Fatal(err)
	}

	d.DB = db
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) InsertShowImage(showID int, image Image) {
	if _, err := d.insertShowImage.Exec(
		image.ID,
		showID,
		image.Type,
		image.Source,
	); err != nil {
		log.Println(err)
	}
}

func (d *DB) InsertSeasonImage(seasonID int, image Image) {
	if _, err := d.insertSeasonImage.Exec(
		image.ID,
		seasonID,
		image.Type,
		image.Source,
	); err != nil {
		log.Println(err)
	}
}

func (d *DB) InsertEpisodeImage(episodeID int, image Image) {
	if _, err := d.insertEpisodeImage.Exec(
		image.ID,
		episodeID,
		image.Type,
		image.Source,
	); err != nil {
		log.Println(err)
	}
}

func (d *DB) UpdateEpisodeDetails(episodeID int, details xstream.Details) {
	if _, err := d.updateEpisodeDetails.Exec(
		details.Year,
		strings.Join(details.SeriesCast, ", "),
		details.NzRating,
		details.NzRatingReason,
		details.NzRatingAdvisories,
		episodeID,
	); err != nil {
		log.Println(err)
	}
}
