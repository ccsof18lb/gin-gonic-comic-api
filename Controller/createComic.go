package controllers

import (
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	getcollection "gincomicapi/Collection"
	database "gincomicapi/Database"
	model "gincomicapi/Model"
	response "gincomicapi/Response"
)

var comicCollection = getcollection.GetCollection(database.ConnectDB())
var validate = validator.New()

func CreateComic() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
		var comic model.Comic
		defer cancel()

		if err := c.BindJSON(&comic); err != nil {
			c.JSON(
				http.StatusBadRequest,response.ComicResponse{
					Status: http.StatusBadRequest, 
					Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&comic); validationErr != nil {
			c.JSON(http.StatusBadRequest, response.ComicResponse{
				Status: http.StatusBadRequest, 
				Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
		}

		newComic := model.Comic{
			ComicId: comic.ComicId,
			Title: comic.Title,
			ReleaseDate: comic.ReleaseDate,
			TrailerLink: comic.TrailerLink,
			Genres: comic.Genres,
			Poster: comic.Poster,
			Backdrops: comic.Backdrops,
			ReviewIds: comic.ReviewIds,
		}

		result, err := comicCollection.InsertOne(ctx,newComic)
		if err != nil {
			c.JSON(http.StatusInternalServerError,response.ComicResponse{
				Status: http.StatusInternalServerError, 
				Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, 
			response.ComicResponse{
				Status: http.StatusCreated, 
				Message: "success", Data: map[string]interface{}{"data": result}})
	}
}