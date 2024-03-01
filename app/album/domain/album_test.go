package domain_test

import (
	"testing"
	"web-wervice/app/album/domain"
)

func TestNewAlbumShouldGenerateDomainEvent(t *testing.T) {
	album := domain.NewAlbum(new(domain.AlbumId), new(domain.AlbumTitle), new(domain.AlbumArtist), new(domain.AlbumPrice))
	aggregate := album.Aggregate
	if len(aggregate.DomainEvents) != 1 {
		t.Fatalf("Expected 1 domain event, got %d", len(aggregate.DomainEvents))
	}
}
