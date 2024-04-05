package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/entity"
	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/kafka"
	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/repository"
)

func sensorMessage() []byte {
	reading := repository.GenerateReading()

	data, err := os.ReadFile("./assets/sensor.json")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}
	var sensor entity.Sensor
	err = json.Unmarshal(data, &sensor)
	if err != nil {
		log.Fatalf("Erro ao decodificar o JSON: %v", err)
	}

	sensor.Nivel = reading

	// sensor := entity.Sensor{
	// 	IdSensor:     "sensor_001",
	// 	Timestamp:    time.Now(),
	// 	TipoPoluente: "PM2.5",
	// 	Nivel:        reading,
	// }
	//
	jsonData, err := json.Marshal(sensor)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	return jsonData
}

func main() {
	kafkaClient := kafka.GenerateKafkaClient(kafka.Producer)

	defer kafkaClient.Producer.Close()

	for {
		msg := sensorMessage()
		kafkaClient.PublishMessage("test_topic", msg)
		time.Sleep(1 * time.Second)
	}
}
