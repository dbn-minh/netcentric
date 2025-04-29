package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Conversation struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserIDs []string           `bson:"user_ids" json:"user_ids"`
	Topic   string             `bson:"topic" json:"topic"`
}
