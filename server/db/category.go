package db

func (d *DB) ShowCategories(showID int) ([]Category, error) {

	rows, err := d.DB.Query(`
		select
			categories.name
		from
			categories,
			show_categories,
			shows
		where 
			categories.id = show_categories.category_id and
			shows.id = show_categories.show_id and
			shows.id = ?
	`, showID)

	if err != nil {
		return nil, err
	}

	categories := make([]Category, 0)

	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		categories = append(categories, Category(category))
	}

	return categories, nil
}

func (d *DB) CategoryShows(categoryID int) ([]Show, error) {

	rows, err := d.DB.Query(`
		select
			shows.id,
			shows.name,
			shows.poster,
			shows.first_aired
		from
			categories,
			show_categories,
			shows
		where 
			categories.id = show_categories.category_id and
			shows.id = show_categories.show_id and
			categories.id = ?
	`, categoryID)

	if err != nil {
		return nil, err
	}

	shows := make([]Show, 0)

	for rows.Next() {
		show := Show{}
		if err := rows.Scan(
			&show.ID,
			&show.Name,
			&show.Poster,
			&show.FirstAired,
		); err != nil {
			return nil, err
		}
		shows = append(shows, show)
	}

	return shows, nil
}
