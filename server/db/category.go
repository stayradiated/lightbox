package db

import (
	"database/sql"
	"strconv"
)

func (d *DB) Categories() (map[string]Category, error) {

	rows, err := d.DB.Query(`
		select
			categories.id,
			categories.name,
			categories.list_id
		from
			categories
		order by
			categories.name
	`)

	if err != nil {
		return nil, err
	}

	categories := make(map[string]Category)

	for rows.Next() {
		var category Category
		var listID sql.NullInt64
		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&listID,
		); err != nil {
			return nil, err
		}
		if listID.Valid {
			category.List = int(listID.Int64)
		} else {
			category.List = -1
		}
		categories[strconv.Itoa(category.ID)] = category
	}

	return categories, nil
}

func (d *DB) ShowCategories(showID int) ([]int, error) {

	rows, err := d.DB.Query(`
		select
			show_categories.category_id
		from
			show_categories,
			shows
		where 
			shows.id = show_categories.show_id and
			shows.id = ?
	`, showID)

	if err != nil {
		return nil, err
	}

	categories := make([]int, 0)

	for rows.Next() {
		var category int
		if err := rows.Scan(
			&category,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (d *DB) CategoryShows(categoryID int) ([]Show, error) {

	rows, err := d.DB.Query(`
		select
			shows.id, shows.title, shows.poster, shows.released
		from
			categories,
			show_categories,
			shows
		where 
			categories.id = show_categories.category_id and
			shows.id = show_categories.show_id and
			categories.id = ?
	`, categoryID)

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
			&show.Released,
		); err != nil {
			return nil, err
		}
		shows = append(shows, show)
	}

	return shows, nil
}
