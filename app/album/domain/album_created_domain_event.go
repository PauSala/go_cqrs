package domain

import (
	"time"
	"web-wervice/app/shared/domain/bus/event"

	"github.com/google/uuid"
)

type AlbumCreatedDomainEvent struct {
	Aggregate_id string
	Dto          interface{}
	Event_id     string
	Date         string
}

func NewAlbumCreatedDomainEvent(a Album) *AlbumCreatedDomainEvent {
	event := AlbumCreatedDomainEvent{}
	event.FromPrimitives(a.AggregateId(), a.ToPrimitives(), uuid.New().String(), time.Now().String())
	return &event
}

func (a *AlbumCreatedDomainEvent) FromPrimitives(aggregate_id string, dto interface{}, event_id string, date string) {
	a.Aggregate_id = aggregate_id
	a.Dto = dto
	a.Event_id = event_id
	a.Date = date
}

func (a *AlbumCreatedDomainEvent) ToPrimitives() event.DomainEventDto {
	return event.DomainEventDto{
		Dto: a.Dto,
	}
}
