package create

import (
	"fmt"
	command "web-wervice/app/shared/domain/bus/command"
)

type CreateAlbumCommandHandler struct {
	AlbumCreator AlbumCreator
}

func (c CreateAlbumCommandHandler) Handle(command command.Command) error {
	createAlbumCommand, ok := command.(*CreateAlbumCommand)
	if !ok {
		return fmt.Errorf("invalid command type expected *CreateAlbumCommand, got %T", command)
	}
	id := createAlbumCommand.Id
	title := createAlbumCommand.Title
	artist := createAlbumCommand.Artist
	price := createAlbumCommand.Price
	return c.AlbumCreator.Run(artist, id, price, title)
}
