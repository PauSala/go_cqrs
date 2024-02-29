package domain

type AlbumRepository interface {
	Save(album Album)
	GetAll() []Album
	GetByID(id string) (Album, error)
}
