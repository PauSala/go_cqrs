package query

type FindAlbumQuery struct {
	id string
}

func (a *FindAlbumQuery) Id() string {
	return a.id
}
