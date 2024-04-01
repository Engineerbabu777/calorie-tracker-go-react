package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	// CONNECTION STRING!
	MongoDB := "conn string..."

	// CREATING A NEW CLIENT!
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))

	// ERROR HANDLING!
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// CONNECTING TO DB!
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client // RETURNING!
}

var Client *mongo.Client = DBInstance();

/*
 	 	CONNECTING TO MONGODB COLLECTION 
		TO PERFORM ACTIONS!
*/
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("caloriesdb").Collection(collectionName);
	return collection;
}