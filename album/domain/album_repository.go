package domain

type AlbumRepository interface {
	Insert(album Album)
	GetAll() []Album
	GetByID(id string) (Album, error)
}
