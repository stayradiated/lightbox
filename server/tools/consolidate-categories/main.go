package main

import (
	"fmt"
	"log"
)

func main() {

	db := new(DB)
	db.Init()
	defer db.Close()

	catsMap := db.GetAllCategories()
	fmt.Println(catsMap)

	for _, id := range db.GetAllShows() {

		cats := db.GetLBShowCategories(id)

		for _, cat := range cats {

			catID, ok := catsMap[cat]

			if !ok {
				log.Fatal("Missing: " + cat)
			}

			db.InsertShowCategory(id, catID)

		}

	}

}
