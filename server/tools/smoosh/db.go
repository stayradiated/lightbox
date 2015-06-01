package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DB *sql.DB

	getAllShows      *sql.Stmt
	insertMasterShow *sql.Stmt

	getAllSeasons      *sql.Stmt
	getMatchingSeason  *sql.Stmt
	insertMasterSeason *sql.Stmt

	getAllEpisodes      *sql.Stmt
	getMatchingEpisode  *sql.Stmt
	insertMasterEpisode *sql.Stmt

	getShowIDs             *sql.Stmt
	getAllCategories       *sql.Stmt
	getIMDBCategories      *sql.Stmt
	getTVDBCategories      *sql.Stmt
	getLBCategories        *sql.Stmt
	insertMasterCategories *sql.Stmt
}

func (d *DB) Init() {
	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox_backup")
	if err != nil {
		log.Fatal(err)
	}

	if d.getAllShows, err = db.Prepare(`
		select
			lb_series.id, lb_series.date_created, lb_series.title, lb_series.description, lb_series.parental_rating,
			shows.id, shows.name, shows.overview, shows.genre, shows.actors, shows.content_rating, shows.first_aired, shows.language, shows.network, shows.rating, shows.rating_count, shows.runtime, shows.status, shows.banner, shows.fanart, shows.poster,
			imdb.id, imdb.title, imdb.year, imdb.rated, imdb.released, imdb.runtime, imdb.genre, imdb.writer, imdb.plot, imdb.language, imdb.country, imdb.poster, imdb.imdb_rating, imdb.imdb_votes, imdb.imdb_id,
			lb_show_images.Source

		from lb_series

		left join shows
			on shows.lightbox_id = lb_series.id

		left join imdb_shows as imdb
			on imdb.lightbox_id = lb_series.id

		left join lb_show_images
			on lb_show_images.show_id = lb_series.id

		group by lb_show_images.show_id
			
	`); err != nil {
		log.Fatal(err)
	}

	if d.insertMasterShow, err = db.Prepare(`
		insert ignore into master_shows (
			id, title, year, released, runtime, writer, actors, plot, poster, fanart,
			rating, rating_count, date_created, parental_rating, imdb, tvdb
		)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	if d.getAllSeasons, err = db.Prepare(`
		select
			lb_seasons.id, lb_seasons.series_id, lb_seasons.date_created, lb_seasons.season_number, lb_seasons.title, lb_seasons.description, lb_seasons.parental_rating,
			lb_season_images.source
		from lb_seasons
		left join lb_season_images on lb_season_images.season_id = lb_seasons.id
		group by lb_seasons.id;
	`); err != nil {
		log.Fatal(err)
	}

	if d.getMatchingSeason, err = db.Prepare(`
		select seasons.id, seasons.show_id, seasons.number, seasons.banner
		from seasons, shows, lb_series, lb_seasons
		where
			seasons.show_id = shows.id and
			shows.lightbox_id = lb_series.id and
			lb_seasons.series_id = lb_series.id and
			lb_seasons.season_number = seasons.number and
			lb_seasons.id = ?;
	`); err != nil {
		log.Fatal(err)
	}

	if d.insertMasterSeason, err = db.Prepare(`
		insert ignore into master_seasons (
			id, show_id, date_created, number, parental_rating, image, tvdb
		)
		values (?, ?, ?, ?, ?, ?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	if d.getAllEpisodes, err = db.Prepare(`
		select
			id, season_id, date_published, date_created, title, description, parental_rating, episode_number, media_id, runtime, air_date, nz_rating, nz_rating_reason, series_cast, year
		from lb_episodes;
	`); err != nil {
		log.Fatal(err)
	}

	if d.getMatchingEpisode, err = db.Prepare(`
		select
			episodes.id,
			episodes.show_id,
			episodes.season_id,
			episodes.director,
			episodes.episode_name,
			episodes.episode_number,
			episodes.first_aired,
			episodes.guest_stars,
			episodes.imdb,
			episodes.language,
			episodes.overview,
			episodes.rating,
			episodes.rating_count,
			episodes.season_number,
			episodes.writer,
			episodes.filename
		from episodes, seasons, shows, lb_episodes, lb_series, lb_seasons
		where
			-- shows = lb_series
			shows.lightbox_id = lb_series.id and
			-- seasons = lb_seasons
			seasons.number = lb_seasons.season_number and
			-- episodes = lb_episodes
			episodes.episode_number = lb_episodes.episode_number and
			-- shows = seasons
			shows.id = seasons.show_id and
			-- seasons = episodes
			seasons.id = episodes.season_id and
			-- lb_series = lb_seasons
			lb_series.id = lb_seasons.series_id and
			-- lb_seasons = lb_episodes
			lb_seasons.id = lb_episodes.season_id and
			-- lb_episode id
			lb_episodes.id = ?;
	`); err != nil {
		log.Fatal(err)
	}

	if d.insertMasterEpisode, err = db.Prepare(`
		insert ignore into master_episodes (
			id,
			season_id,
			media_id,
			date_created,
			date_published,
			number,
			title,
			plot,
			runtime,
			first_aired,
			year,
			parental_rating,
			parental_rating_reason,
			director,
			writer,
			guest_stars,
			rating,
			rating_count,
			image,
			imdb,
			tvdb
		)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	if d.getShowIDs, err = db.Prepare(`
		select id from lb_series;
	`); err != nil {
		log.Fatal(err)
	}

	if d.getAllCategories, err = db.Prepare(`
		select id, name
		from master_categories
	`); err != nil {
		log.Fatal(err)
	}

	if d.getIMDBCategories, err = db.Prepare(`
		select master_categories.name
		from master_categories, imdb_show_categories, imdb_shows
		where 
			master_categories.id = imdb_show_categories.category_id and
			imdb_shows.id = imdb_show_categories.show_id and
			imdb_shows.lightbox_id = ?;
	`); err != nil {
		log.Fatal(err)
	}

	if d.getTVDBCategories, err = db.Prepare(`
		select master_categories.name
		from master_categories, show_categories, shows
		where 
			master_categories.id = show_categories.category_id and
			shows.id = show_categories.show_id and
			shows.lightbox_id = ?;
	`); err != nil {
		log.Fatal(err)
	}

	if d.getLBCategories, err = db.Prepare(`
		select lb_categories.title
		from lb_categories, lb_series_categories
		where 
			lb_categories.id = lb_series_categories.category_id and
			lb_series_categories.series_id = ?;
	`); err != nil {
		log.Fatal(err)
	}

	if d.insertMasterCategories, err = db.Prepare(`
		insert ignore into master_show_categories(
			show_id, category_id
		) values(?, ?)
	`); err != nil {
		log.Fatal(err)
	}

	d.DB = db
}

func (d *DB) Close() {
	d.DB.Close()
}

func (d *DB) GetAllShows() []ShowData {
	rows, err := d.getAllShows.Query()
	if err != nil {
		log.Fatal(err)
	}

	showData := make([]ShowData, 0)

	for rows.Next() {
		var s ShowData
		if err := rows.Scan(
			&s.ID,
			&s.LB.DateCreated,
			&s.LB.Title,
			&s.LB.Description,
			&s.LB.ParentalRating,
			&s.TVDB.ID,
			&s.TVDB.Name,
			&s.TVDB.Overview,
			&s.TVDB.Genre,
			&s.TVDB.Actors,
			&s.TVDB.ContentRating,
			&s.TVDB.FirstAired,
			&s.TVDB.Language,
			&s.TVDB.Network,
			&s.TVDB.Rating,
			&s.TVDB.RatingCount,
			&s.TVDB.Runtime,
			&s.TVDB.Status,
			&s.TVDB.Banner,
			&s.TVDB.Fanart,
			&s.TVDB.Poster,
			&s.IMDB.ID,
			&s.IMDB.Title,
			&s.IMDB.Year,
			&s.IMDB.Rated,
			&s.IMDB.Released,
			&s.IMDB.Runtime,
			&s.IMDB.Genre,
			&s.IMDB.Writer,
			&s.IMDB.Plot,
			&s.IMDB.Language,
			&s.IMDB.Country,
			&s.IMDB.Poster,
			&s.IMDB.Rating,
			&s.IMDB.Votes,
			&s.IMDB.IMDB,
			&s.LB.Image,
		); err != nil {
			log.Fatal(err)
		}
		showData = append(showData, s)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return showData
}

func (d *DB) InsertMasterShow(show Show) {
	if _, err := d.insertMasterShow.Exec(
		show.ID,
		show.Title,
		show.Year,
		show.Released,
		show.Runtime,
		show.Writer,
		show.Actors,
		show.Plot,
		show.Poster,
		show.Fanart,
		show.Rating,
		show.RatingCount,
		show.DateCreated,
		show.ParentalRating,
		show.IMDB,
		show.TVDB,
	); err != nil {
		log.Fatal(err)
	}
}

func (d *DB) GetAllSeasons() []SeasonData {

	rows, err := d.getAllSeasons.Query()
	if err != nil {
		log.Fatal(err)
	}

	seasonData := make([]SeasonData, 0)

	for rows.Next() {
		var s SeasonData
		if err := rows.Scan(
			&s.ID,
			&s.LB.ShowID,
			&s.LB.DateCreated,
			&s.LB.Number,
			&s.LB.Title,
			&s.LB.Description,
			&s.LB.ParentalRating,
			&s.LB.Image,
		); err != nil {
			log.Fatal(err)
		}

		d.getMatchingSeason.QueryRow(s.ID).Scan(
			&s.TVDB.ID,
			&s.TVDB.ShowID,
			&s.TVDB.Number,
			&s.TVDB.Banner,
		)

		seasonData = append(seasonData, s)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return seasonData
}

func (d *DB) InsertMasterSeason(season Season) {
	if _, err := d.insertMasterSeason.Exec(
		season.ID,
		season.ShowID,
		season.DateCreated,
		season.Number,
		season.ParentalRating,
		season.Image,
		season.TVDB,
	); err != nil {
		j, _ := json.Marshal(season)
		fmt.Println(string(j))
		log.Fatal(err)
	}
}

func (d *DB) GetAllEpisodes() []EpisodeData {

	rows, err := d.getAllEpisodes.Query()
	if err != nil {
		log.Fatal(err)
	}

	episodeData := make([]EpisodeData, 0)

	for rows.Next() {
		var e EpisodeData

		if err := rows.Scan(
			&e.ID,
			&e.LB.SeasonID,
			&e.LB.DatePublished,
			&e.LB.DateCreated,
			&e.LB.Title,
			&e.LB.Description,
			&e.LB.ParentalRating,
			&e.LB.Number,
			&e.LB.MediaID,
			&e.LB.Runtime,
			&e.LB.AirDate,
			&e.LB.NZRating,
			&e.LB.NZRatingReason,
			&e.LB.SeriesCast,
			&e.LB.Year,
		); err != nil {
			log.Fatal(err)
		}

		d.getMatchingEpisode.QueryRow(e.ID).Scan(
			&e.TVDB.ID,
			&e.TVDB.ShowID,
			&e.TVDB.SeasonID,
			&e.TVDB.Director,
			&e.TVDB.Name,
			&e.TVDB.Number,
			&e.TVDB.FirstAired,
			&e.TVDB.GuestStars,
			&e.TVDB.IMDB,
			&e.TVDB.Language,
			&e.TVDB.Overview,
			&e.TVDB.Rating,
			&e.TVDB.RatingCount,
			&e.TVDB.SeasonNumber,
			&e.TVDB.Writer,
			&e.TVDB.Filename,
		)

		episodeData = append(episodeData, e)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return episodeData
}

func (d *DB) InserMasterEpisode(episode Episode) {
	if _, err := d.insertMasterEpisode.Exec(
		episode.ID,
		episode.SeasonID,
		episode.MediaID,
		episode.DateCreated,
		episode.DatePublished,
		episode.Number,
		episode.Title,
		episode.Plot,
		episode.Runtime,
		episode.FirstAired,
		episode.Year,
		episode.ParentalRating,
		episode.ParentalRatingReason,
		episode.Director,
		episode.Writer,
		episode.GuestStars,
		episode.Rating,
		episode.RatingCount,
		episode.Image,
		episode.IMDB,
		episode.TVDB,
	); err != nil {
		j, _ := json.Marshal(episode)
		fmt.Println(string(j))
		log.Fatal(err)
	}
}

func (d *DB) GetShowIDs() []int {
	return queryIntSlice(d.getShowIDs)
}

func (d *DB) GetShowCategories(showID int) CategoryData {
	categoryData := CategoryData{
		IMDB: queryStringSlice(d.getIMDBCategories, showID),
		TVDB: queryStringSlice(d.getTVDBCategories, showID),
		LB:   queryStringSlice(d.getLBCategories, showID),
	}

	return categoryData
}

func (d *DB) GetAllCategories() map[string]int {
	rows, err := d.getAllCategories.Query()
	if err != nil {
		log.Fatal(err)
	}

	cats := make(map[string]int, 0)

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		cats[name] = id
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return cats
}

func (d *DB) InsertMasterCategories(showID, categoryID int) {
	if _, err := d.insertMasterCategories.Exec(
		showID, categoryID,
	); err != nil {
		log.Fatal(err)
	}
}

func queryIntSlice(stmt *sql.Stmt, data ...interface{}) []int {
	rows, err := stmt.Query(data...)
	if err != nil {
		log.Fatal(err)
	}

	arr := make([]int, 0)

	for rows.Next() {
		var n int
		if err := rows.Scan(&n); err != nil {
			log.Fatal(err)
		}
		arr = append(arr, n)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}

func queryStringSlice(stmt *sql.Stmt, data ...interface{}) []string {
	rows, err := stmt.Query(data...)
	if err != nil {
		log.Fatal(err)
	}

	arr := make([]string, 0)

	for rows.Next() {
		var n string
		if err := rows.Scan(&n); err != nil {
			log.Fatal(err)
		}
		arr = append(arr, n)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return arr
}
