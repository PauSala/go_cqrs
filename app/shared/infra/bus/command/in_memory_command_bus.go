package command

import (
	"errors"
	"web-wervice/app/shared/domain/bus/command"
)

type InMemoryCommandBus struct {
	handlers map[string]command.CommandHandler
}

func (b *InMemoryCommandBus) RegisterHandler(c command.Command, h command.CommandHandler) {
	b.handlers[c.CommandId()] = h
}

func (b *InMemoryCommandBus) Dispatch(c command.Command) error {
	handler := b.handlers[c.CommandId()]
	if handler != nil {
		return handler.Handle(c)
	}
	//TODO: type errors
	return errors.New("no handler for command: " + c.CommandId())
}

func CreateInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers: make(map[string]command.CommandHandler),
	}
}
