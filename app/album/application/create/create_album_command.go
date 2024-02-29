package create

type CreateAlbumCommand struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func NewCreateAlbumCommand(id string, title string, artist string, price float64) CreateAlbumCommand {
	return CreateAlbumCommand{Id: id, Title: title, Artist: artist, Price: price}
}

func (a *CreateAlbumCommand) CommandId() string {
	return "CreateAlbumCommand"
}
