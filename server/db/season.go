package db

func (d *DB) ShowSeasons(showID int) ([]Season, error) {

	rows, err := d.DB.Query(`
		select
			id,
			show_id,
			number,
			banner
		from
			seasons
		where
			show_id = (?)
	`, showID)

	if err != nil {
		return nil, err
	}

	seasons := make([]Season, 0)

	for rows.Next() {
		var season Season
		if err := rows.Scan(
			&season.ID,
			&season.ShowID,
			&season.Number,
			&season.Banner,
		); err != nil {
			return nil, err
		}
		seasons = append(seasons, season)
	}

	return seasons, nil
}

func (d *DB) Season(seasonID int) (Season, error) {

	season := Season{}

	if err := d.DB.QueryRow(`
		select
			id,
			show_id,
			number,
			banner
		from
			seasons
		where
			id = (?)
	`, seasonID).Scan(
		&season.ID,
		&season.ShowID,
		&season.Number,
		&season.Banner,
	); err != nil {
		return season, err
	}

	episodes, err := d.SeasonEpisodes(seasonID)
	if err != nil {
		return season, err
	}
	season.Episodes = episodes

	return season, nil
}
