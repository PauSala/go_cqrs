package aggregate

type AggregateRoot[T any] interface {
	PulldomainEvents() []T
	Record(domainEvent T)
}

type Aggregate[T any] struct {
	DomainEvents []T
}

func (a *Aggregate[T]) Record(domainEvent T) {
	a.DomainEvents = append(a.DomainEvents, domainEvent)
}

func (a *Aggregate[T]) PulldomainEvents() []T {
	domainEvents := a.DomainEvents
	a.DomainEvents = []T{}
	return domainEvents
}
