package command

type CommandHandler interface {
	Handle(c Command) error
}
