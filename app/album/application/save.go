package application

import "web-wervice/app/album/domain"

type SaveAlbumService struct {
	AlbumRepository domain.AlbumRepository
}

func (r *SaveAlbumService) Run(album domain.Album) {
	r.AlbumRepository.Insert(album)
}
