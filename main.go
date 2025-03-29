package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanover/gohelloworld/album"
)


func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
	router.GET("/albums", album.GetAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)
	router.POST("/albums", album.PostAlbums)

	router.Run("localhost:8080")
}