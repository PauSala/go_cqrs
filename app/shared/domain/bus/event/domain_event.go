package event

type DomainEventDto struct {
	Dto interface{}
}

type DomainEvent interface {
	FromPrimitives(aggregate_id string, dto interface{}, event_id string, date string)
	ToPrimitives() DomainEventDto
}
