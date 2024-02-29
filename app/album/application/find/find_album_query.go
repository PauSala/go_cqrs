package find

type FindAlbumQuery struct {
	id string
}

func (a FindAlbumQuery) Id() string {
	return a.id
}

func (a FindAlbumQuery) QueryId() string {
	return "FindAlbumQuery"
}

func NewFindAlbumQuery(id string) FindAlbumQuery {
	return FindAlbumQuery{id: id}
}
