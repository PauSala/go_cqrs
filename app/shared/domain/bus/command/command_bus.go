package command

type CommandBus interface {
	Dispatch(command Command) error
	Publish(event any) error
}
