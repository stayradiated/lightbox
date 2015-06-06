package db

func (d *DB) Lists() ([]List, error) {

	rows, err := d.DB.Query(`
		select
			lists.id,
			lists.title
		from
			lists
	`)

	if err != nil {
		return nil, err
	}

	lists := make([]List, 0)

	for rows.Next() {
		var list List
		if err := rows.Scan(
			&list.ID, &list.Title,
		); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (d *DB) ListShows(listID int) ([]int, error) {

	rows, err := d.DB.Query(`
		select
			list_shows.show_id
		from
			list_shows
		where
			list_shows.list_id = ?
		order by
			list_shows.n
	`, listID)

	if err != nil {
		return nil, err
	}

	shows := make([]int, 0)

	for rows.Next() {
		var showID int
		if err := rows.Scan(
			&showID,
		); err != nil {
			return nil, err
		}
		shows = append(shows, showID)
	}

	return shows, nil
}
