package controllers

import (
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	model "gincomicapi/Model"
	response "gincomicapi/Response"
	"go.mongodb.org/mongo-driver/bson"
)

func EditOneComic() gin.HandlerFunc {
	return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        comicId := c.Param("comicId")
        var comic model.Comic
        defer cancel()

        //validate the request body
        if err := c.BindJSON(&comic); err != nil {
            c.JSON(http.StatusBadRequest, response.ComicResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&comic); validationErr != nil {
            c.JSON(http.StatusBadRequest, response.ComicResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        update := bson.M{
			"title": comic.Title,
			"releaseDate": comic.ReleaseDate,
			"trailerLink": comic.TrailerLink,
			"genres": comic.Genres,
			"poster": comic.Poster,
			"backdrops": comic.Backdrops,
			"reviewIds": comic.ReviewIds,
		}
        result, err := comicCollection.UpdateOne(ctx, bson.M{"comicId": comicId}, bson.M{"$set": update})
        if err != nil {
            c.JSON(http.StatusInternalServerError, response.ComicResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        var updatedComic model.Comic
        if result.MatchedCount == 1 {
            err := comicCollection.FindOne(ctx, bson.M{"comicId": comicId}).Decode(&updatedComic)
            if err != nil {
                c.JSON(http.StatusInternalServerError, response.ComicResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                return
            }
        }

        c.JSON(http.StatusOK, response.ComicResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedComic}})
    }
}