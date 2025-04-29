package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"netcentric-nosql/config"
	"netcentric-nosql/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var convCollection *mongo.Collection

func InitConversationCollection() {
	convCollection = config.DB.Collection("conversations")
}

func CreateConversation(w http.ResponseWriter, r *http.Request) {
	var conv models.Conversation
	json.NewDecoder(r.Body).Decode(&conv)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, _ := convCollection.InsertOne(ctx, conv)
	json.NewEncoder(w).Encode(res)
}

func GetConversations(w http.ResponseWriter, r *http.Request) {
	var convs []models.Conversation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, _ := convCollection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var conv models.Conversation
		cursor.Decode(&conv)
		convs = append(convs, conv)
	}
	json.NewEncoder(w).Encode(convs)
}
