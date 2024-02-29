package domain

type AlbumArtist struct {
	value string
}

func NewAlbumArtist(value string) *AlbumArtist {
	return &AlbumArtist{value: value}
}

func (id *AlbumArtist) Value() string {
	return id.value
}

func (id *AlbumArtist) Equals(other string) bool {
	return id.value == other
}
