package infra

import (
	"web-wervice/app/album/application"
	"web-wervice/app/album/application/create"
	"web-wervice/app/album/domain"
)

var albums = []domain.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type AlbumServiceFactory struct {
	GetAlbumsService    application.GetAlbumsService
	SaveAlbumService    create.AlbumCreator
	GetAllAlbumsService application.GetAllAlbumsService
}

func NewAlbumServiceFactory() AlbumServiceFactory {
	albumRepository := NewInMemoryAlbumRepository(albums)
	return AlbumServiceFactory{
		GetAlbumsService:    application.GetAlbumsService{AlbumRepository: albumRepository},
		SaveAlbumService:    create.AlbumCreator{AlbumRepository: albumRepository},
		GetAllAlbumsService: application.GetAllAlbumsService{AlbumRepository: albumRepository},
	}
}
