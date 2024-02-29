package infra

import (
	"web-wervice/app/album/application/create"
	"web-wervice/app/album/application/find"
	"web-wervice/app/album/domain"
	"web-wervice/app/shared/domain/bus/command"
)

var albums = []domain.Album{
	{ID: domain.NewAlbumId("1"), Title: domain.NewAlbumTitle("Blue Train"), Artist: domain.NewAlbumArtist("John Coltrane"), Price: domain.NewAlbumPrice(32.3)},
	{ID: domain.NewAlbumId("2"), Title: domain.NewAlbumTitle("Jeru"), Artist: domain.NewAlbumArtist("Gerry Mulligan"), Price: domain.NewAlbumPrice(12.13)},
	{ID: domain.NewAlbumId("3"), Title: domain.NewAlbumTitle("Sarah Vaughan and Clifford Brown"), Artist: domain.NewAlbumArtist("Sarah Vaughan"), Price: domain.NewAlbumPrice(22.3)},
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
