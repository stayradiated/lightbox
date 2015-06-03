package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DB *sql.DB

	getShowImages    *sql.Stmt
	getShowFanart    *sql.Stmt
	getSeasonImages  *sql.Stmt
	getEpisodeImages *sql.Stmt
}

func (d *DB) Init() {
	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}

	if d.getShowImages, err = db.Prepare(`
		select id, poster
		from shows
		where poster != ""
	`); err != nil {
		log.Fatal(err)
	}

	if d.getShowFanart, err = db.Prepare(`
		select id, fanart
		from shows
		where fanart != ""
	`); err != nil {
		log.Fatal(err)
	}

	if d.getSeasonImages, err = db.Prepare(`
		select id, image
		from seasons
		where image != ""
	`); err != nil {
		log.Fatal(err)
	}

	if d.getEpisodeImages, err = db.Prepare(`
		select id, image
		from episodes
		where image != ""
	`); err != nil {
		log.Fatal(err)
	}

	d.DB = db
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) GetAllImages() ImageData {
	return ImageData{
		Shows:    queryImages(d.getShowImages),
		Fanart:   queryImages(d.getShowFanart),
		Seasons:  queryImages(d.getSeasonImages),
		Episodes: queryImages(d.getEpisodeImages),
	}
}

func queryImages(stmt *sql.Stmt) []Image {
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	images := make([]Image, 0)

	for rows.Next() {
		var image Image
		if err := rows.Scan(&image.ID, &image.Source); err != nil {
			log.Fatal(err)
		}
		images = append(images, image)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return images
}
