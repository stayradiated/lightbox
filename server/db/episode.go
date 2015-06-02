package db

func (d *DB) Episode(episodeID int) (Episode, error) {

	episode := Episode{}

	if err := d.DB.QueryRow(`
		select
			episodes.id,
			episodes.season_id,
			seasons.show_id,
			episodes.media_id,
			episodes.date_created,
			episodes.date_published,
			episodes.number,
			episodes.title,
			episodes.plot,
			episodes.runtime,
			episodes.first_aired,
			episodes.year,
			episodes.parental_rating,
			episodes.parental_rating_reason,
			episodes.director,
			episodes.writer,
			episodes.guest_stars,
			episodes.rating,
			episodes.rating_count,
			episodes.image,
			episodes.imdb,
			episodes.tvdb
		from
			episodes,
			seasons
		where
			episodes.id = (?) and
			episodes.season_id = seasons.id
	`, episodeID).Scan(
		&episode.ID,
		&episode.SeasonID,
		&episode.ShowID,
		&episode.MediaID,
		&episode.DateCreated,
		&episode.DatePublished,
		&episode.Number,
		&episode.Title,
		&episode.Plot,
		&episode.Runtime,
		&episode.FirstAired,
		&episode.Year,
		&episode.ParentalRating,
		&episode.ParentalRatingReason,
		&episode.Director,
		&episode.Writer,
		&episode.GuestStars,
		&episode.Rating,
		&episode.RatingCount,
		&episode.Image,
		&episode.IMDB,
		&episode.TVDB,
	); err != nil {
		return episode, err
	}

	return episode, nil
}

func (d *DB) SeasonEpisodes(seasonID int) ([]Episode, error) {

	rows, err := d.DB.Query(`
		select
			episodes.id,
			episodes.season_id,
			seasons.show_id,
			episodes.media_id,
			episodes.date_created,
			episodes.date_published,
			episodes.number,
			episodes.title,
			episodes.plot,
			episodes.runtime,
			episodes.first_aired,
			episodes.year,
			episodes.parental_rating,
			episodes.parental_rating_reason,
			episodes.director,
			episodes.writer,
			episodes.guest_stars,
			episodes.rating,
			episodes.rating_count,
			episodes.image,
			episodes.imdb,
			episodes.tvdb
		from
			episodes,
			seasons
		where
			episodes.season_id = seasons.id and
			season_id = ?
		order by
			number
	`, seasonID)

	if err != nil {
		return nil, err
	}

	episodes := make([]Episode, 0)

	for rows.Next() {
		episode := Episode{}
		if err = rows.Scan(
			&episode.ID,
			&episode.SeasonID,
			&episode.ShowID,
			&episode.MediaID,
			&episode.DateCreated,
			&episode.DatePublished,
			&episode.Number,
			&episode.Title,
			&episode.Plot,
			&episode.Runtime,
			&episode.FirstAired,
			&episode.Year,
			&episode.ParentalRating,
			&episode.ParentalRatingReason,
			&episode.Director,
			&episode.Writer,
			&episode.GuestStars,
			&episode.Rating,
			&episode.RatingCount,
			&episode.Image,
			&episode.IMDB,
			&episode.TVDB,
		); err != nil {
			return nil, err
		}
		episodes = append(episodes, episode)
	}

	return episodes, nil
}
