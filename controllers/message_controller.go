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

var messageCollection *mongo.Collection

func InitMessageCollection() {
	messageCollection = config.DB.Collection("messages")
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	json.NewDecoder(r.Body).Decode(&message)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, _ := messageCollection.InsertOne(ctx, message)
	json.NewEncoder(w).Encode(res)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []models.Message
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, _ := messageCollection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var msg models.Message
		cursor.Decode(&msg)
		messages = append(messages, msg)
	}
	json.NewEncoder(w).Encode(messages)
}
