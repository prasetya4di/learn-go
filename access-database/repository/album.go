package repository

import "data-access/model"

type AlbumRepository interface {
	Read() ([]model.Album, error)
	ReadByArtist(name string) ([]model.Album, error)
	ReadById(id int64) (model.Album, error)
	Insert(album model.Album) (int64, error)
}
