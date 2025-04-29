package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderID       string             `bson:"sender_id" json:"sender_id"`
	ConversationID string             `bson:"conversation_id" json:"conversation_id"`
	Content        string             `bson:"content" json:"content"`
}
