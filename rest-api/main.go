package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"web-service-gin/model"
)

func main() {
	router := gin.Default()
	router.GET("/albums", model.GetAlbums)
	router.GET("/albums/:id", model.GetAlbumId)
	router.POST("/albums", model.PostAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
