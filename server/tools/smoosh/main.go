package main

func main() {

	db := new(DB)
	db.Init()
	defer db.Close()

	db.SmooshShows()
	db.SmooshSeasons()
	db.SmooshEpisodes()
	// db.SmooshCategories()

}
