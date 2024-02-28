package application

import "web-wervice/album/domain"

type SaveAlbumService struct {
	AlbumRepository domain.AlbumRepository
}

func (r *SaveAlbumService) Run(album domain.Album) {
	r.AlbumRepository.Insert(album)
}
