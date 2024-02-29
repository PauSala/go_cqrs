package query

import "web-wervice/app/album/domain"

type AlbumFinder struct {
	AlbumRepository domain.AlbumRepository
}

func (r *AlbumFinder) Run(id string) (domain.Album, error) {
	return r.AlbumRepository.GetByID(id)
}
