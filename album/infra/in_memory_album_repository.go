package infra

import "web-wervice/album/domain"

type InMemoryAlbumRepository struct {
	albums []domain.Album
}

func (r *InMemoryAlbumRepository) Insert(album domain.Album) {
	r.albums = append(r.albums, album)
}

func (r *InMemoryAlbumRepository) GetAll() []domain.Album {
	return r.albums
}

type AlbumNotFound struct{}

func (m *AlbumNotFound) Error() string {
	return "AlbumNotFound"
}

func (r *InMemoryAlbumRepository) GetByID(id string) (domain.Album, error) {
	for _, album := range r.albums {
		if album.ID == id {
			return album, nil
		}
	}
	return domain.Album{}, &AlbumNotFound{} // return zero value if not found
}

func NewInMemoryAlbumRepository(initialAlbums []domain.Album) domain.AlbumRepository {
	return &InMemoryAlbumRepository{albums: initialAlbums}
}
