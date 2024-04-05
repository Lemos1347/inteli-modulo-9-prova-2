package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Message string             `bson:"message"`
}
