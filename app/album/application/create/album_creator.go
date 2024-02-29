package create

import (
	"web-wervice/app/album/domain"
	"web-wervice/app/shared/domain/bus/command"
)

type AlbumCreator struct {
	AlbumRepository domain.AlbumRepository
	command.CommandBus
}

func (r *AlbumCreator) Run(artist string, id string, price float64, title string) error {
	Artist := domain.NewAlbumArtist(artist)
	Id := domain.NewAlbumId(id)
	Price := domain.NewAlbumPrice(price)
	Title := domain.NewAlbumTitle(title)
	album := domain.NewAlbum(Id, Title, Artist, Price)
	r.AlbumRepository.Save(album)
	var events = album.Aggregate.PulldomainEvents()
	r.CommandBus.Publish(events)
	return nil
}

func NewAlbumCreator(r domain.AlbumRepository) *AlbumCreator {
	return &AlbumCreator{AlbumRepository: r}
}
