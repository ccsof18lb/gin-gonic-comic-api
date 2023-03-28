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

func GetOneComic() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        comicId := c.Param("comicId")
        var comic model.Comic
        defer cancel()

        err := comicCollection.FindOne(ctx, bson.M{"comicId": comicId}).Decode(&comic)
        if err != nil {
            c.JSON(http.StatusInternalServerError, 
				response.ComicResponse{Status: http.StatusInternalServerError, 
					Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }
		c.JSON(http.StatusOK, 
			response.ComicResponse{Status: http.StatusOK, 
				Message: "success", Data: map[string]interface{}{"data": comic}})
	}
}