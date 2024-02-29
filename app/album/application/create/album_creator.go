package create

import "web-wervice/app/album/domain"

type AlbumCreator struct {
	AlbumRepository domain.AlbumRepository
}

func (r *AlbumCreator) Run(album domain.Album) error {
	r.AlbumRepository.Insert(album)
	return nil
}
