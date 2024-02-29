package application

import "web-wervice/app/album/domain"

type GetAlbumsService struct {
	AlbumRepository domain.AlbumRepository
}

func (r *GetAlbumsService) Run(id string) (domain.Album, error) {
	return r.AlbumRepository.GetByID(id)
}
