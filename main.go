package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	controllers "gincomicapi/Controller"
)

func main() {
	fmt.Println("Server running")
	router := gin.Default()

	router.Use(cors.Default())
	
	router.GET("/comics",controllers.ReadAllComic())

	router.POST("/new-comic", controllers.CreateComic())

	router.GET("/comics/:comicId", controllers.GetOneComic())

	router.PUT("/update-comic/:comicId", controllers.EditOneComic())

	router.Run("localhost:5000")
}