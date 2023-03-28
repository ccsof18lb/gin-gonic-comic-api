package getcollection

import (
 "go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("johnjohn").Collection("maymay")
	return collection
}