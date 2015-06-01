package db

func (d *DB) Shows(filter string, limit, offset int) ([]Show, error) {

	filter = "%" + filter + "%"

	if offset < 0 {
		offset = 0
	}

	if limit == 0 {
		limit = 24
	}

	// if limit > 50 {
	// 	limit = 50
	// }

	rows, err := d.DB.Query(`
		select
			id, title, poster, year, released, date_created
		from
			shows
		where
			title like (?)
		order by title asc
		limit ?
		offset ?
	`, filter, limit, offset)

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
			&show.Year,
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
