package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	mongoURI := "mongodb://localhost:27017"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	dbName := "proba"
	collectionName := "roles"

	err = client.Database(dbName).Collection(collectionName).Drop(context.Background())
	if err != nil {
		log.Println("Failed to drop collection:", err)
	}

	collection := client.Database(dbName).Collection(collectionName)

	document := bson.M{
		"name1": "",
		"name2": "",
	}

	_, err = collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration completed successfully.")
}
