package domain

type AlbumId struct {
	value string
}

func NewAlbumId(value string) *AlbumId {
	return &AlbumId{value: value}
}

func (id *AlbumId) Value() string {
	return id.value
}

func (id *AlbumId) Equals(other string) bool {
	return id.value == other
}
