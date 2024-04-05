package main

import (
	"encoding/json"
	"log"

	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/db"
	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/entity"
	kafkaI "github.com/Lemos1347/inteli-modulo-9-prova-2/internal/kafka"
	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/repository"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func handleKafkaMessage(msg *kafka.Message, dbClient db.DBClient) {
	// fmt.Printf("Received message: %s\n", string(msg.Value))

	var sensor entity.Sensor

	json.Unmarshal(msg.Value, &sensor)

	log.Printf("\n --------- \n \033[1;32mSensor data received from sensor %s:\033[0m \n -timestamp: %s \n -tipoPoluente: %s \n -nivel: %d \n --------- \n", sensor.IdSensor, sensor.Timestamp, sensor.TipoPoluente, sensor.Nivel)

	repository.CreateSensor(&sensor, dbClient)
}

func main() {
	dbClient := db.NewDBConnection()

	callbackFunc := func(msg *kafka.Message) {
		handleKafkaMessage(msg, *dbClient)
	}

	kafkaClient := kafkaI.GenerateKafkaClient(kafkaI.Consumer)

	defer kafkaClient.Consumer.Close()

	kafkaClient.StartConsumingMessages("test_topic", callbackFunc)
}
