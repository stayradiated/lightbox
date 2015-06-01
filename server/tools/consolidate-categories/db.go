package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DB *sql.DB

	getAllShows         *sql.Stmt
	getAllCategories    *sql.Stmt
	getLBShowCategories *sql.Stmt

	insertShowCategory *sql.Stmt
}

func (d *DB) Init() {
	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}

	if d.getAllShows, err = db.Prepare(`
		select lb_series.id
		from lb_series
	`); err != nil {
		log.Fatal(err)
	}

	if d.getAllCategories, err = db.Prepare(`
		select id, name
		from categories
	`); err != nil {
		log.Fatal(err)
	}

	if d.getLBShowCategories, err = db.Prepare(`
		select lb_categories.title
		from lb_categories, lb_series_categories, lb_series
		where
			lb_series.id = lb_series_categories.series_id and
			lb_categories.id = lb_series_categories.category_id and 
			lb_series.id = ?
	`); err != nil {
		log.Fatal(err)
	}

	if d.insertShowCategory, err = db.Prepare(`
		insert ignore into lb_cat (show_id, cat_id)
		values (?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	d.DB = db
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) GetAllCategories() map[string]int {
	rows, err := d.getAllCategories.Query()
	if err != nil {
		log.Fatal(err)
	}

	cats := make(map[string]int, 0)

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		cats[name] = id
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return cats
}

func (d *DB) GetAllShows() []int {
	rows, err := d.getAllShows.Query()
	if err != nil {
		log.Fatal(err)
	}

	ids := make([]int, 0)

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return ids
}

func (d *DB) GetLBShowCategories(showID int) []string {
	rows, err := d.getLBShowCategories.Query(showID)
	if err != nil {
		log.Fatal(err)
	}

	cats := make([]string, 0)

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			log.Fatal(err)
		}
		cats = append(cats, s)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return cats
}

func (d *DB) InsertShowCategory(showID, catID int) {
	if _, err := d.insertShowCategory.Exec(showID, catID); err != nil {
		log.Fatal(err)
	}
}
