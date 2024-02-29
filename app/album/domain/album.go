package domain

import (
	"web-wervice/app/shared/domain/aggregate"
	"web-wervice/app/shared/domain/bus/event"
	"web-wervice/app/shared/domain/valueObject"
)

type Album struct {
	ID        valueObject.ValueObject[string]
	Title     valueObject.ValueObject[string]
	Artist    valueObject.ValueObject[string]
	Price     valueObject.ValueObject[float64]
	Aggregate aggregate.Aggregate
}

type AlbumDto struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

func (a *Album) ToPrimitives() AlbumDto {
	return AlbumDto{
		ID:     a.ID.Value(),
		Title:  a.Title.Value(),
		Artist: a.Artist.Value(),
		Price:  a.Price.Value(),
	}
}

func (a *Album) AggregateId() string {
	return "album"
}

func NewAlbum(id *AlbumId, title *AlbumTitle, artist *AlbumArtist, price *AlbumPrice) Album {
	album := Album{
		ID:        id,
		Title:     title,
		Artist:    artist,
		Price:     price,
		Aggregate: aggregate.Aggregate{DomainEvents: make([]event.DomainEvent, 0)},
	}
	event := NewAlbumCreatedDomainEvent(album)
	album.Aggregate.Record(event)
	return album
}
