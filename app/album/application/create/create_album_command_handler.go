package create

import "web-wervice/app/album/domain"

type CreateAlbumCommandHandler struct {
	AlbumCreator AlbumCreator
}

func (c *CreateAlbumCommandHandler) Handle(command *CreateAlbumCommand) error {
	id := command.Id()
	title := command.Tittle()
	artist := command.Artist()
	price := command.Price()
	album := domain.Album{Artist: artist, ID: id, Price: price, Title: title}
	return c.AlbumCreator.Run(album)
}
