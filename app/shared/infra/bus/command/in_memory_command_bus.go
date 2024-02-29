package command

import (
	"errors"
	"web-wervice/app/shared/domain/bus/command"
)

type InMemoryCommandBus struct {
	handlers map[string]command.CommandHandler
}

func (b *InMemoryCommandBus) RegisterHandler(c command.Command, h command.CommandHandler) {
	b.handlers[c.Id()] = h
}

func CreateInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers: make(map[string]command.CommandHandler),
	}
}

func (b *InMemoryCommandBus) Dispatch(c command.Command) error {
	handler := b.handlers[c.Id()]
	if handler != nil {
		handler.Handle(c)
	}
	//TODO: type errors
	return errors.New("no handler for command: " + c.Id())
}
