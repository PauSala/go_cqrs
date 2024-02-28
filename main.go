package main

import (
	"net/http"
	album "web-wervice/album/domain"
	infra "web-wervice/album/infra"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

var factory = infra.NewAlbumServiceFactory()

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	service := factory.GetAllAlbumsService
	albums := service.Run()
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an Album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	service := factory.SaveAlbumService
	service.Run(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the Album whose ID value matches the id
// parameter sent by the client, then returns that Album as a response.
func getAlbumByID(c *gin.Context) {

	id := c.Param("id")
	service := factory.GetAlbumsService
	album, err := service.Run(id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
