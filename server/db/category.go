package db

func (d *DB) Categories() ([]Category, error) {

	rows, err := d.DB.Query(`
		select
			categories.id,
			categories.name
		from
			categories
		order by
			categories.name
	`)

	if err != nil {
		return nil, err
	}

	categories := make([]Category, 0)

	for rows.Next() {
		var category Category
		if err := rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (d *DB) ShowCategories(showID int) ([]Category, error) {

	rows, err := d.DB.Query(`
		select
			categories.id,
			categories.name
		from
			categories,
			show_categories,
			shows
		where 
			categories.id = show_categories.category_id and
			shows.id = show_categories.show_id and
			shows.id = ?
		order by
			categories.name
	`, showID)

	if err != nil {
		return nil, err
	}

	categories := make([]Category, 0)

	for rows.Next() {
		var category Category
		if err := rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (d *DB) CategoryShows(categoryID int) ([]Show, error) {

	rows, err := d.DB.Query(`
		select
			shows.id, shows.title, shows.poster, shows.released
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
			&show.Title,
			&show.Poster,
			&show.Released,
		); err != nil {
			return nil, err
		}
		shows = append(shows, show)
	}

	return shows, nil
}
