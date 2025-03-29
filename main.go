package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanover/gohelloworld/album"
)

func main() {
	router := gin.Default()
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/api/albums", album.GetAlbums)
	router.GET("/api/albums/:id", album.GetAlbumByID)
	router.POST("/api/albums", album.PostAlbums)

	router.Run(":2112")
}
