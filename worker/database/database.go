package database

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewDatabase(url string, database string) *mongo.Database {
	connectOption := options.Client().ApplyURI(url)
	client, err := mongo.Connect(connectOption)
	if err != nil {
		log.Fatalf("MongoDB client create error: %v", err)
	}
	return client.Database(database)
}
