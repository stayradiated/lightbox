package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var category int

	flag.IntVar(&category, "c", -1, "Category")
	flag.Parse()

	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		select id, title 
		from lb_series, lb_series_categories
		where
			lb_series.id = lb_series_categories.series_id
			and category_id = ?
	`, category)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id    int
		title string
	)

	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("-id=%d -name=\"%s\"\n", id, title)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
