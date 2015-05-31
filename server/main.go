package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/stayradiated/lightbox/server/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// setup database
	database, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// create router
	router := NewRouter(&Handlers{&db.DB{database}})

	// start http server
	fmt.Println("Starting server on :9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
