package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://akshatsrivastava:3feb2004chotu@akshat.kfjhgxh.mongodb.net/"
const dbname = "Education_Portal"

// MongoDB client
var client *mongo.Client

// Init initializes the MongoDB connection and client
func init() {
	// Set client options
	clientoptions := options.Client().ApplyURI(connectionstring)

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientoptions)
	if err != nil {
		panic(err)
	}

	fmt.Println("MONGODB CONNECTION")

	// Test the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
}

// GetCollection returns a MongoDB collection for the given model name
func GetCollection(collectionName string) *mongo.Collection {
	return client.Database(dbname).Collection(collectionName)
}
