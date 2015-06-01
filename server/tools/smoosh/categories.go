package main

import "log"

type CategoryData struct {
	IMDB []string
	TVDB []string
	LB   []string
}

func (d *DB) SmooshCategories() {

	catIDs := d.GetAllCategories()

	// get all shows in master
	for _, showID := range d.GetShowIDs() {

		categories := d.GetShowCategories(showID)

		set := make(map[string]bool)
		for _, name := range categories.IMDB {
			set[name] = true
		}
		for _, name := range categories.TVDB {
			set[name] = true
		}
		for _, name := range categories.LB {
			set[name] = true
		}

		for name := range set {

			catID, ok := catIDs[name]

			if !ok {
				log.Fatal("MISSING CATEGORY", name)
			}

			d.InsertMasterCategories(showID, catID)

		}

	}

}
