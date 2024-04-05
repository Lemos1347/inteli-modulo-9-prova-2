package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sensor struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	IdSensor     string             `bson:"idSensor"`
	Timestamp    time.Time          `bson:"timestamp"`
	TipoPoluente string             `bson:"tipoPoluente"`
	Nivel        int                `bson:"nivel"`
}
