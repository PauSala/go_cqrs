package main

import (
	"net/http"
	create "web-wervice/app/album/application/create"
	find "web-wervice/app/album/application/find"
	infra "web-wervice/app/album/infra"
	command "web-wervice/app/shared/domain/bus/command"
	query "web-wervice/app/shared/domain/bus/query"
	commandShared "web-wervice/app/shared/infra/bus/command"
	queryShared "web-wervice/app/shared/infra/bus/query"

	"github.com/gin-gonic/gin"
)

var commandBus = commandShared.CreateInMemoryCommandBus()
var queryBus = queryShared.CreateInMemoryQueryBus()
var factory = infra.NewAlbumServiceFactory(commandBus)

func main() {
	var command_handler command.CommandHandler = create.CreateAlbumCommandHandler{AlbumCreator: factory.AlbumCreator}
	commandBus.RegisterHandler(&create.CreateAlbumCommand{}, command_handler)

	var query_handler query.QueryHandler = find.FindAlbumQueryHandler{AlbumFinder: factory.AlbumFinder}
	queryBus.RegisterHandler(&find.FindAlbumQuery{}, query_handler)

	router := gin.Default()
	router.GET("/albums/:id", albumFinderController)
	router.POST("/albums", postAlbumController)
	router.Run("localhost:8080")
}

// postAlbumController adds an Album from JSON received in the request body.
func postAlbumController(c *gin.Context) {
	var newAlbum create.CreateAlbumCommand

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	command := create.NewCreateAlbumCommand(newAlbum.Id, newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	res := commandBus.Dispatch(&command)
	if res != nil {
		println("Error creating album" + res.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating album"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Album created"})
}

// albumFinderController locates the Album whose ID value matches the id
// parameter sent by the client, then returns that Album as a response.
func albumFinderController(c *gin.Context) {

	id := c.Param("id")

	query := find.NewFindAlbumQuery(id)
	album, err := queryBus.Ask(&query)
	if err == nil {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
