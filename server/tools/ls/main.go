package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, title from series")
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
