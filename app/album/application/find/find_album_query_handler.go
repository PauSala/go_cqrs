package find

import (
	"fmt"
	"web-wervice/app/album/domain"
	"web-wervice/app/shared/domain/bus/query"
)

type FindAlbumQueryHandler struct {
	AlbumFinder AlbumFinder
}

func (h FindAlbumQueryHandler) Handle(query query.Query) (interface{}, error) {
	_query, ok := query.(*FindAlbumQuery)
	if !ok {
		return domain.Album{}, fmt.Errorf("invalid command type expected *FindAlbumQuery, got %T", _query)
	}
	id := _query.Id()
	return h.AlbumFinder.Run(id)
}
