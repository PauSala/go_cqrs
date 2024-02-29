package query

import "web-wervice/app/album/domain"

type FindAlbumQueryHandler struct {
	AlbumFinder AlbumFinder
}

func (h *FindAlbumQueryHandler) Handle(query *FindAlbumQuery) (domain.Album, error) {
	id := query.Id()
	return h.AlbumFinder.Run(id)
}
