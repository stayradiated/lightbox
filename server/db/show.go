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
			id, name, poster, first_aired
		from
			shows
		where
			name like (?)
		order by name asc
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

func (d *DB) Show(showID int) (Show, error) {

	show := Show{}

	if err := d.DB.QueryRow(`
		select
			id,
			name,
			overview,
			rating,
			rating_count,
			actors,
			poster,
			fanart,
			content_rating,
			first_aired,
			runtime,
			imdb
		from
			shows
		where
			id = (?)
	`, showID).Scan(
		&show.ID,
		&show.Name,
		&show.Overview,
		&show.Rating,
		&show.RatingCount,
		&show.Actors,
		&show.Poster,
		&show.Fanart,
		&show.ContentRating,
		&show.FirstAired,
		&show.Runtime,
		&show.IMDB,
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
