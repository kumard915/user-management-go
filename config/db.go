package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://dummy:Kumard915@cluster0.ohsrlfg.mongodb.net/mydatabase?retryWrites=true&w=majority&appName=Cluster0")

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	// Select the database
	DB = client.Database("mydatabase")
	log.Println("Connected to MongoDB")
}
