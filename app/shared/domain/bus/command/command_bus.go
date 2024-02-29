package command

import (
	"web-wervice/app/shared/domain/bus/event"
)

type CommandBus interface {
	Dispatch(command Command) error
	Publish(event []event.DomainEvent) error
}
