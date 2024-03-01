package create_test

import (
	"testing"
	"web-wervice/app/album/application/create"
	"web-wervice/app/album/domain"
	"web-wervice/app/album/infra"
	"web-wervice/app/shared/domain/bus/command"
)

func TestAlbumCreatorShouldPublishAnEventWhenCalled(t *testing.T) {
	albumRepository := infra.NewInMemoryAlbumRepository([]domain.Album{})
	commandBus := &command.MockCommandBus{}
	albumCreator := create.NewAlbumCreator(albumRepository, commandBus)
	albumCreator.Run(new(
		domain.AlbumArtist).Value(),
		new(domain.AlbumId).Value(),
		new(domain.AlbumPrice).Value(),
		new(domain.AlbumTitle).Value())
	if commandBus.PublishHasBeenCalled != 1 {
		t.Fatalf("Expected command bus to have been called")
	}
}
