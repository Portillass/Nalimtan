package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record albums.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// In-memory data (temporary database)
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "Rell", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "portillas", Price: 17.99},
	{ID: "3", Title: "lol found hahh", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	// Routes
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", createAlbum)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}

// GET all albums
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// GET album by ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// CREATE album
func createAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// Validation
	if newAlbum.ID == "" || newAlbum.Title == "" || newAlbum.Artist == "" || newAlbum.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "all fields are required and price must be greater than 0s",
		})
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// UPDATE album
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			// Validation
			if updatedAlbum.Title == "" || updatedAlbum.Artist == "" || updatedAlbum.Price <= 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "title, artist and price are required",
				})
				return
			}

			albums[i].Title = updatedAlbum.Title
			albums[i].Artist = updatedAlbum.Artist
			albums[i].Price = updatedAlbum.Price

			c.JSON(http.StatusOK, albums[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// DELETE album
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
	c.JSON()
}

