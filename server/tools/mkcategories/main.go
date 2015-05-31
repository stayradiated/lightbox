package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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

	rows, err := db.Query(`
		select id, name from categories
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	categories := make(map[string]int)

	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		categories[name] = id
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	insertSeriesCategory, err := db.Prepare(`
		insert ignore into show_categories(show_id, category_id) values (?, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query(`
		select id, genre from shows
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id    int
		genre string
	)

	for rows.Next() {
		if err := rows.Scan(&id, &genre); err != nil {
			log.Fatal(err)
		}

		genre = strings.Trim(genre, "|")

		for _, category := range strings.Split(genre, "|") {
			catID, ok := categories[category]
			if ok {
				insertSeriesCategory.Exec(id, catID)
			} else {
				fmt.Println("Missing Category", category)
			}
		}

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
