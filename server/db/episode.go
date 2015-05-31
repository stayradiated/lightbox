package db

func (d *DB) Episode(episodeID int) (Episode, error) {

	episode := Episode{}

	if err := d.DB.QueryRow(`
		select
			id,
			episode_number,
			episode_name,
			overview,
			first_aired,
			filename,
			rating,
			rating_count,
			director,
			writer,
			guest_stars,
			imdb
		from
			episodes
		where
			id = (?)
	`, episodeID).Scan(
		&episode.ID,
		&episode.Number,
		&episode.Name,
		&episode.Overview,
		&episode.FirstAired,
		&episode.Image,
		&episode.Rating,
		&episode.RatingCount,
		&episode.Director,
		&episode.Writer,
		&episode.GuestStars,
		&episode.IMDB,
	); err != nil {
		return episode, err
	}

	return episode, nil
}

func (d *DB) SeasonEpisodes(seasonID int) ([]Episode, error) {

	rows, err := d.DB.Query(`
		select
			id,
			episode_number,
			episode_name,
			overview,
			first_aired,
			filename,
			rating,
			rating_count,
			director,
			writer,
			guest_stars,
			imdb
		from
			episodes
		where
			season_id = (?)
	`, seasonID)

	if err != nil {
		return nil, err
	}

	episodes := make([]Episode, 0)

	for rows.Next() {
		episode := Episode{}
		if err = rows.Scan(
			&episode.ID,
			&episode.Number,
			&episode.Name,
			&episode.Overview,
			&episode.FirstAired,
			&episode.Image,
			&episode.Rating,
			&episode.RatingCount,
			&episode.Director,
			&episode.Writer,
			&episode.GuestStars,
			&episode.IMDB,
		); err != nil {
			return nil, err
		}
		episodes = append(episodes, episode)
	}

	return episodes, nil
}
