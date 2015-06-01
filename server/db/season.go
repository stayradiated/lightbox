package db

func (d *DB) ShowSeasons(showID int) ([]Season, error) {

	rows, err := d.DB.Query(`
		select
			seasons.id,
			seasons.show_id,
			seasons.date_created,
			seasons.number,
			seasons.parental_rating,
			seasons.image,
			seasons.tvdb,
			count(episodes.id) as episode_count
		from
			seasons, episodes
		where
			episodes.season_id = seasons.id and
			seasons.show_id = ?
		group by
			seasons.id
		order by
			seasons.number
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
			&season.DateCreated,
			&season.Number,
			&season.ParentalRating,
			&season.Image,
			&season.TVDB,
			&season.EpisodeCount,
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
			seasons.id,
			seasons.show_id,
			seasons.date_created,
			seasons.number,
			seasons.parental_rating,
			seasons.image,
			seasons.tvdb,
			count(episodes.id) as episode_count
		from
			seasons, episodes
		where
			episodes.season_id = seasons.id and
			seasons.id = ?
	`, seasonID).Scan(
		&season.ID,
		&season.ShowID,
		&season.DateCreated,
		&season.Number,
		&season.ParentalRating,
		&season.Image,
		&season.TVDB,
		&season.EpisodeCount,
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
