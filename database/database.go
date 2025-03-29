package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var AlbumsCollection *mongo.Collection

// ConnectDB establishes a connection to MongoDB
func ConnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace with your MongoDB connection string
	clientOptions := options.Client().ApplyURI("mongodb://tom:password123@gotest-mongodb-kcuixf:27017")
	var err error

	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = Client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// Initialize collection
	AlbumsCollection = Client.Database("musicdb").Collection("albums")

	log.Println("Connected to MongoDB")
	return nil
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB")
}
