package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Bue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "John Coltrane", Price: 19.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 32.14},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsById(c *gin.Context) {
	var id string = c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "El álbum no fue encontrado!",
	})
}

func setAlbum(c *gin.Context) {
	var id string = c.Param("id")
	var newAlbumInfo album
	if err := c.BindJSON(&newAlbumInfo); err != nil {
		return
	}

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			albums = append(albums, newAlbumInfo)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "El álbum no fue encontrado!",
	})
}

func deleteAlbum(c *gin.Context) {
	var id string = c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "El álbum no fue encontrado!",
	})
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)
	router.PUT("/album/:id", setAlbum)
	router.DELETE("/album/:id", deleteAlbum)

	router.Run("localhost:8080")
}
