package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thanover/gohelloworld/album"
	"github.com/thanover/gohelloworld/database"
)

func main() {
	// Connect to MongoDB
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer database.DisconnectDB()

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
