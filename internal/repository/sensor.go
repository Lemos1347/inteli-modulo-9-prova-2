package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/db"
	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateReading() int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	num := r.Intn(62) + 15
	return num
}

func CreateSensor(sensor *entity.Sensor, db db.DBClient) error {
	result, err := db.DB.Collection("sensors").InsertOne(context.TODO(), sensor)
	if err != nil {
		errMsg := fmt.Sprintf("Error creating sensor: %s\n", err.Error())
		panic(errMsg)
	}

	sensor.Id = result.InsertedID.(primitive.ObjectID)

	fmt.Printf("Message created: %#v\n", *sensor)

	return nil
}
