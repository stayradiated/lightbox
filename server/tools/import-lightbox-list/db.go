package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DB *sql.DB

	insertList     *sql.Stmt
	insertListShow *sql.Stmt
}

func (d *DB) Init() {
	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}

	if d.insertList, err = db.Prepare(`
		insert ignore into lists (
			id, title
		)
		values (?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	if d.insertListShow, err = db.Prepare(`
		insert ignore into list_shows (
			list_id, show_id, n
		)
		values (?, ?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	d.DB = db
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) InsertList(showID int, title string) {
	if _, err := d.insertList.Exec(
		showID, title,
	); err != nil {
		log.Println(err)
	}
}

func (d *DB) InsertListShow(listID, showID, order int) {
	if _, err := d.insertListShow.Exec(
		listID, showID, order,
	); err != nil {
		log.Println(err)
	}
}
