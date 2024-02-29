package create

type CreateAlbumCommand struct {
	id     string
	title  string
	artist string
	price  float64
}

func (a *CreateAlbumCommand) CommandId() string {
	return "CreateAlbumCommand"
}

func (a *CreateAlbumCommand) Id() string {
	return a.id
}

func (a *CreateAlbumCommand) Tittle() string {
	return a.title
}

func (a *CreateAlbumCommand) Artist() string {
	return a.artist
}

func (a *CreateAlbumCommand) Price() float64 {
	return a.price
}
