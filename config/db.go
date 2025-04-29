package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	fmt.Println("🔌 Connecting to MongoDB...")

	clientOptions := options.Client().ApplyURI("mongodb+srv://nhatminhak2003:Doanminh2003.@cluster0.crki1.mongodb.net/")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	DB = client.Database("chat-app_netcentric")
	fmt.Println("✅ MongoDB connection established.")
}
