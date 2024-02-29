package query

import (
	"errors"
	"web-wervice/app/shared/domain/bus/query"
)

type InMemoryQueryBus struct {
	handlers map[string]query.QueryHandler
}

func (b *InMemoryQueryBus) RegisterHandler(c query.Query, h query.QueryHandler) {
	b.handlers[c.QueryId()] = h
}

func (b *InMemoryQueryBus) Ask(c query.Query) (interface{}, error) {
	handler := b.handlers[c.QueryId()]
	if handler != nil {
		return handler.Handle(c)
	}
	return nil, errors.New("no handler for command: " + c.QueryId())
}

func CreateInMemoryCommandBus() *InMemoryQueryBus {
	return &InMemoryQueryBus{
		handlers: make(map[string]query.QueryHandler),
	}
}
