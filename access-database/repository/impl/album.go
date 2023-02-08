package impl

import (
	"data-access/model"
	"data-access/repository"
	"database/sql"
	"fmt"
)

type albumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) repository.AlbumRepository {
	return albumRepository{
		db: db,
	}
}

func (ar albumRepository) Read() ([]model.Album, error) {
	var albums []model.Album

	rows, err := ar.db.Query("Select * from album")
	if err != nil {
		return nil, fmt.Errorf("Get album failed : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb model.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("Get album failed : %v", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Get album failed : %v", err)
	}
	return albums, nil
}

func (ar albumRepository) ReadByArtist(name string) ([]model.Album, error) {
	var albums []model.Album

	rows, err := ar.db.Query("Select * from album where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb model.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func (ar albumRepository) ReadById(id int64) (model.Album, error) {
	// An album to hold data from the returned row
	var alb model.Album

	row := ar.db.QueryRow("SELECT * from album where id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func (ar albumRepository) Insert(album model.Album) (int64, error) {
	result, err := ar.db.Exec("Insert into album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
