package infra

import (
	"web-wervice/app/album/application/create"
	"web-wervice/app/album/application/find"
	"web-wervice/app/album/domain"
	"web-wervice/app/shared/domain/bus/command"
)

var albums = []domain.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type AlbumServiceFactory struct {
	AlbumFinder  find.AlbumFinder
	AlbumCreator create.AlbumCreator
}

func NewAlbumServiceFactory(commandBus command.CommandBus) AlbumServiceFactory {
	albumRepository := NewInMemoryAlbumRepository(albums)
	return AlbumServiceFactory{
		AlbumFinder:  find.AlbumFinder{AlbumRepository: albumRepository},
		AlbumCreator: create.AlbumCreator{AlbumRepository: albumRepository, CommandBus: commandBus},
	}
}
