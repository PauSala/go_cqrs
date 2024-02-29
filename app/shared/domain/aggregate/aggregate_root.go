package aggregate

import (
	"web-wervice/app/shared/domain/bus/event"
)

type AggregateRoot interface {
	PulldomainEvents() []event.DomainEvent
	Record(domainEvent event.DomainEvent)
}

type Aggregate struct {
	DomainEvents []event.DomainEvent
}

func (a *Aggregate) Record(domainEvent event.DomainEvent) {
	a.DomainEvents = append(a.DomainEvents, domainEvent)
}

func (a *Aggregate) PulldomainEvents() []event.DomainEvent {
	domainEvents := a.DomainEvents
	a.DomainEvents = []event.DomainEvent{}
	return domainEvents
}
