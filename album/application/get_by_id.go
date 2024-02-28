package application

import "web-wervice/album/domain"

type GetAlbumsService struct {
	AlbumRepository domain.AlbumRepository
}

func (r *GetAlbumsService) Run(id string) (domain.Album, error) {
	return r.AlbumRepository.GetByID(id)
}
