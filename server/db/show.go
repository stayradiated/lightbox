package db

func (d *DB) Shows(filter string) ([]Show, error) {

	filter = "%" + filter + "%"

	rows, err := d.DB.Query(`
		select
			id, title, year, rating, parental_rating, released, date_created
		from
			shows
		where
			title like (?)
		order by
			rating desc
	`, filter)

	if err != nil {
		return nil, err
	}

	shows := make([]Show, 0)

	for rows.Next() {
		show := Show{}
		if err := rows.Scan(
			&show.ID,
			&show.Title,
			&show.Year,
			&show.Rating,
			&show.ParentalRating,
			&show.Released,
			&show.DateCreated,
		); err != nil {
			return nil, err
		}
		shows = append(shows, show)
	}

	return shows, nil
}

func (d *DB) Show(showID int) (Show, error) {

	show := Show{}

	if err := d.DB.QueryRow(`
		select
			id,
			title,
			year,
			released,
			runtime,
			writer,
			actors,
			plot,
			poster,
			fanart,
			rating,
			rating_count,
			date_created,
			parental_rating,
			imdb,
			tvdb
		from
			shows
		where
			id = (?)
	`, showID).Scan(
		&show.ID,
		&show.Title,
		&show.Year,
		&show.Released,
		&show.Runtime,
		&show.Writer,
		&show.Actors,
		&show.Plot,
		&show.Poster,
		&show.Fanart,
		&show.Rating,
		&show.RatingCount,
		&show.DateCreated,
		&show.ParentalRating,
		&show.IMDB,
		&show.TVDB,
	); err != nil {
		return show, err
	}

	categories, err := d.ShowCategories(showID)
	if err != nil {
		return show, err
	}
	show.Categories = categories

	seasons, err := d.ShowSeasons(showID)
	if err != nil {
		return show, err
	}
	show.Seasons = seasons

	return show, nil
}
