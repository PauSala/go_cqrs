package domain

type AlbumTitle struct {
	value string
}

func NewAlbumTitle(value string) *AlbumTitle {
	return &AlbumTitle{value: value}
}

func (id *AlbumTitle) Value() string {
	return id.value
}

func (id *AlbumTitle) Equals(other string) bool {
	return id.value == other
}
