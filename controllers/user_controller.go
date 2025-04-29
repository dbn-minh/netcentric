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

var userCollection *mongo.Collection

func InitUserCollection() {
	userCollection = config.DB.Collection("users")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, _ := userCollection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(res)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, _ := userCollection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}
