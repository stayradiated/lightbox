package main

import (
	"encoding/json"
	"fmt"
	"os"

	"bitbucket.org/stayradiated/lightbox/server/xstream"
)

func main() {

	db := new(DB)
	db.Init()
	defer db.Close()

	var list xstream.List
	d := json.NewDecoder(os.Stdin)
	if err := d.Decode(&list); err != nil {
		panic(err)
	}

	db.InsertList(list.ID, list.Titles.Default)

	for i, series := range list.Elements.Series {
		fmt.Println(series.ID)
		db.InsertListShow(list.ID, series.ID, i)
	}

}
