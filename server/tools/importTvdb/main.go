package main

import (
	"database/sql"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"bitbucket.org/stayradiated/lightbox/server/tvdb"

	_ "github.com/go-sql-driver/mysql"
)

var getID = regexp.MustCompile(`^\d+`)

func main() {
	if err := ImportAllShows("../fetchtvdb/series"); err != nil {
		panic(err)
	}
}

func ImportAllShows(dir string) error {

	// load shows from series folder
	contents, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// connect to mysql
	db, err := sql.Open("mysql",
		"lightbox:lightbox@tcp(192.168.1.100:3306)/lightbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// prepare commands
	insertSeries, err := db.Prepare(`INSERT IGNORE INTO series_tvdb(
		id,
		lightbox_id,
		actors,
		airs_dayofweek,
		airs_time,
		content_rating,
		first_aired,
		genre,
		imdb,
		language,
		network,
		overview,
		rating,
		rating_count,
		runtime,
		series_id,
		series_name,
		status,
		banner,
		fanart,
		lastupdate,
		poster
	) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	insertEpisode, err := db.Prepare(`INSERT IGNORE INTO episodes_tvdb(
		id,
		combined_episodenumber,
		combined_season,
		director,
		episode_name,
		episode_number,
		first_aired,
		guest_stars,
		imdb,
		language,
		overview,
		rating,
		rating_count,
		season_number,
		writer,
		absolute_number,
		filename,
		lastupdated,
		season_id,
		series_id
	) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	insertBanner, err := db.Prepare(`INSERT IGNORE INTO banners(
		id,
		series_id,
		banner_path,
		banner_type,
		banner_size,
		colors,
		language,
		rating,
		rating_count,
		season,
		series_name,
		thumbnail_path,
		vignette_path
	) VALUES(?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range contents {
		if fi.IsDir() == false {
			continue
		}

		name := fi.Name()

		lightboxID, err := strconv.Atoi(getID.FindString(name))
		if err != nil {
			return err
		}

		info := tvdb.SeriesInfo{}
		DecodeXMLFile(&info, filepath.Join(dir, name, "en.xml"))

		banners := tvdb.Banners{}
		DecodeXMLFile(&banners, filepath.Join(dir, name, "banners.xml"))

		series := info.Series

		if _, err := insertSeries.Exec(
			series.ID,
			lightboxID,
			series.Actors,
			series.AirsDayOfWeek,
			series.AirsTime,
			series.ContentRating,
			series.FirstAired,
			series.Genre,
			series.IMDB,
			series.Language,
			series.Network,
			series.Overview,
			series.Rating,
			series.RatingCount,
			series.Runtime,
			series.SeriesID,
			series.SeriesName,
			series.Status,
			series.Banner,
			series.FanArt,
			series.LastUpdated,
			series.Poster,
		); err != nil {
			return err
		}

		for _, banner := range banners.Banners {

			if _, err := insertBanner.Exec(
				banner.ID,
				series.ID,
				banner.BannerPath,
				banner.BannerType,
				banner.BannerSize,
				banner.Colors,
				banner.Language,
				banner.Rating,
				banner.RatingCount,
				banner.Season,
				banner.SeriesName,
				banner.ThumbnailPath,
				banner.VignettePath,
			); err != nil {
				return err
			}

		}

		for _, ep := range info.Episodes {

			if _, err := insertEpisode.Exec(
				ep.ID,
				ep.CombinedEpisodeNumber,
				ep.CombinedSeason,
				ep.Director,
				ep.EpisodeName,
				ep.EpisodeNumber,
				ep.FirstAired,
				ep.GuestStars,
				ep.IMDB,
				ep.Language,
				ep.Overview,
				ep.Rating,
				ep.RatingCount,
				ep.SeasonNumber,
				ep.Writer,
				ep.AbsoluteNumber,
				ep.FileName,
				ep.LastUpdated,
				ep.SeasonID,
				ep.SeriesID,
			); err != nil {
				return err
			}

		}

	}

	return nil
}

func DecodeXMLFile(v interface{}, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	x := xml.NewDecoder(f)
	err = x.Decode(v)
	if err != nil {
		return err
	}

	return nil
}
