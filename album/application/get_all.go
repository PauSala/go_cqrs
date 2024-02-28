package application

import "web-wervice/album/domain"

type GetAllAlbumsService struct {
	AlbumRepository domain.AlbumRepository
}

func (r *GetAllAlbumsService) Run() []domain.Album {
	return r.AlbumRepository.GetAll()
}
