package controllers

import (
	"context"
	"net/http"
	model "gincomicapi/Model"
	response "gincomicapi/Response"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// func enableCors(w *http.ResponseWriter) {
//     (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
// }

func ReadAllComic(
    // w http.ResponseWriter, 
    // c *gin.Context
    ) gin.HandlerFunc {
	return func(c *gin.Context) {
        // enableCors(&w)
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var comics []model.Comic
        defer cancel()

        results, err := comicCollection.Find(ctx, bson.M{})

        if err != nil {
            c.JSON(http.StatusInternalServerError, response.ComicResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //reading from the db in an optimal way
        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleComic model.Comic
            if err = results.Decode(&singleComic); err != nil {
                c.JSON(http.StatusInternalServerError, response.ComicResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }

            comics = append(comics, singleComic)
        }

        c.JSON(http.StatusOK,
            response.ComicResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": comics}},
        )
    }
}