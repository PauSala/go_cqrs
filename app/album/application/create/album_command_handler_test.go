package create_test

import (
	"testing"
	"web-wervice/app/album/application/create"
	"web-wervice/app/album/domain"
	"web-wervice/app/album/infra"
	"web-wervice/app/shared/domain/bus/command"
)

type MockInvalidCommand struct{}

func (m MockInvalidCommand) CommandId() string {
	return "mock_invalid_command"
}

func TestCreateAlbumCommandHandlerShouldRejectInvalidCommands(t *testing.T) {
	albumRepository := infra.NewInMemoryAlbumRepository([]domain.Album{})
	commandBus := &command.MockCommandBus{}
	albumCreator := create.NewAlbumCreator(albumRepository, commandBus)
	albumCreator.Run(new(
		domain.AlbumArtist).Value(),
		new(domain.AlbumId).Value(),
		new(domain.AlbumPrice).Value(),
		new(domain.AlbumTitle).Value())
	commandHandler := create.CreateAlbumCommandHandler{AlbumCreator: *albumCreator}
	err := commandHandler.Handle(&MockInvalidCommand{})
	if err == nil {
		t.Fatalf("Expected error")
	}
}

func TestCreateAlbumCommandHandlerShouldProcessValidCommands(t *testing.T) {
	albumRepository := infra.NewInMemoryAlbumRepository([]domain.Album{})
	commandBus := &command.MockCommandBus{}
	albumCreator := create.NewAlbumCreator(albumRepository, commandBus)
	albumCreator.Run(new(
		domain.AlbumArtist).Value(),
		new(domain.AlbumId).Value(),
		new(domain.AlbumPrice).Value(),
		new(domain.AlbumTitle).Value())
	commandHandler := create.CreateAlbumCommandHandler{AlbumCreator: *albumCreator}
	command := create.NewCreateAlbumCommand("", "", "", 0)
	err := commandHandler.Handle(&command)
	if err != nil {
		t.Fatalf("Expected error, got nil")
	}
}
