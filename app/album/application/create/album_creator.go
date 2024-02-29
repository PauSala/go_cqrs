package create

import (
	"web-wervice/app/album/domain"
	aggregate "web-wervice/app/shared/domain/aggregate"
	"web-wervice/app/shared/domain/bus/command"
)

type AlbumCreator struct {
	AlbumRepository domain.AlbumRepository
	command.CommandBus
}

func (r *AlbumCreator) Run(artist string, id string, price float64, title string) error {
	album := domain.Album{
		Artist:    artist,
		ID:        id,
		Price:     price,
		Title:     title,
		Aggregate: aggregate.Aggregate[domain.Album]{DomainEvents: make([]domain.Album, 0)}}
	r.AlbumRepository.Save(album)
	r.CommandBus.Publish(album)
	return nil
}

func NewAlbumCreator(r domain.AlbumRepository) *AlbumCreator {
	return &AlbumCreator{AlbumRepository: r}
}
