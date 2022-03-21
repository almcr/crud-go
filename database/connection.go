package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client          *mongo.Client
	CrudDb          *mongo.Database
	UsersCollection *mongo.Collection
)

func Init() {
	mongodbUrl := os.Getenv("MONGO_URL")

	log.Println("Connecting to: " + mongodbUrl)
	Client, err := mongo.NewClient(options.Client().ApplyURI(mongodbUrl))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Create a default test db
	CrudDb = Client.Database("test")
	// Create data collection
	UsersCollection = CrudDb.Collection("users_data")

	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB !!")
}
