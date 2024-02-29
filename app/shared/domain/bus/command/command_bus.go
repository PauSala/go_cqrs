package command

type CommandBus interface {
	Dispatch(command Command)
}
