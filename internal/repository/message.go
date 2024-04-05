package repository

import (
	"context"
	"fmt"

	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/db"
	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMessage(message *entity.Message, db db.DBClient) error {
	result, err := db.DB.Collection("messages").InsertOne(context.TODO(), message)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating sensor: %s\n", err.Error())
		panic(errMsg)
	}

	message.Id = result.InsertedID.(primitive.ObjectID)

	fmt.Printf("Message created: %#v\n", *message)

	return nil
}
