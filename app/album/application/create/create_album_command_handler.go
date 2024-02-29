package create

import (
	"fmt"
	"web-wervice/app/album/domain"
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
	id := createAlbumCommand.Id()
	title := createAlbumCommand.Tittle()
	artist := createAlbumCommand.Artist()
	price := createAlbumCommand.Price()
	album := domain.Album{Artist: artist, ID: id, Price: price, Title: title}
	return c.AlbumCreator.Run(album)
}
