package domain

type AlbumPrice struct {
	value float64
}

func NewAlbumPrice(value float64) *AlbumPrice {
	return &AlbumPrice{value: value}
}

func (id *AlbumPrice) Value() float64 {
	return id.value
}

func (id *AlbumPrice) Equals(other float64) bool {
	return id.value == other
}
