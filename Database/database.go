package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetURL() string {
	env_err := godotenv.Load()
	if env_err != nil {log.Fatal(env_err)}
	link := os.Getenv("LINK")
	return link
}

var (
	urll = GetURL()
)

func ConnectDB() *mongo.Client {
	Mongo_URL := urll
	client, err:= mongo.NewClient(options.Client().ApplyURI(Mongo_URL))
	if err != nil {log.Fatal(err)}
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	err = client.Connect(ctx)
	defer cancel()
	if err!= nil {log.Fatal(err)}
	fmt.Println(
		"Connected to the DB")
	return client
}