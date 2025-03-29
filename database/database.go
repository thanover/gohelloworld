package database

import (
	"context"
	"log"
	"os"
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

	// Get MongoDB URI from environment variable
	mongoURI := os.Getenv("MONGO_URI")

	// Use default connection string if environment variable is not set
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
		log.Println("Warning: MONGO_URI environment variable not set, using default connection string")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
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
