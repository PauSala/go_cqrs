package command

import "web-wervice/app/shared/domain/bus/event"

type MockCommandBus struct {
	DispatchHasBeenCalled int32
	PublishHasBeenCalled  int32
}

func (m *MockCommandBus) Dispatch(command Command) error {
	m.DispatchHasBeenCalled += 1
	return nil
}

func (m *MockCommandBus) Publish(event []event.DomainEvent) error {
	m.PublishHasBeenCalled += 1
	return nil
}
