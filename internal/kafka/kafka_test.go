package kafka

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/entity"
	"github.com/Lemos1347/inteli-modulo-9-prova-2/internal/repository"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func generateSensorMessage() []byte {
	reading := repository.GenerateReading()

	sensor := entity.Sensor{
		IdSensor:     "sensor_001",
		Timestamp:    time.Now(),
		TipoPoluente: "PM2.5",
		Nivel:        reading,
	}

	jsonData, err := json.Marshal(sensor)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	return jsonData
}

func TestKafkaSendedMessage(t *testing.T) {
	consumer := GenerateKafkaClient(Consumer)
	producer := GenerateKafkaClient(Producer)

	// defer consumer.Consumer.Close()
	defer producer.Producer.Close()

	message := generateSensorMessage()

	callback := func(msg *kafka.Message) {
		t.Logf("Received message: %s\n", string(msg.Value))

		var sensor entity.Sensor
		var sensorSent entity.Sensor

		json.Unmarshal(msg.Value, &sensor)
		json.Unmarshal(message, &sensorSent)

		t.Logf("Sensor received: %#v\n", sensor)
		t.Logf("Sensor sent: %#v\n", sensorSent)

		if sensor.IdSensor != sensorSent.IdSensor || sensor.Timestamp != sensorSent.Timestamp || sensor.TipoPoluente != sensorSent.TipoPoluente || sensor.Nivel != sensorSent.Nivel {
			t.Fatal("Messages are different!")
		}

	}

	producer.PublishMessage("prova_test", message)

	consumer.StartConsumingMessages("prova_test", callback, true)
}

func TestKafkaMessagePersistence(t *testing.T) {
	consumer1 := GenerateKafkaClient(Consumer)
	consumer2 := GenerateKafkaClient(Consumer, &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"group.id":          "go-consumer-group-2",
		"auto.offset.reset": "earliest",
	})
	producer := GenerateKafkaClient(Producer)

	// defer consumer.Consumer.Close()
	defer producer.Producer.Close()

	message := generateSensorMessage()

	messagesReceived := [2]*entity.Sensor{}

	callback1 := func(msg *kafka.Message) {
		t.Logf("Received message: %s\n", string(msg.Value))

		var sensor entity.Sensor

		json.Unmarshal(msg.Value, &sensor)

		t.Logf("Sensor received: %#v\n", sensor)

		messagesReceived[0] = &sensor

	}

	callback2 := func(msg *kafka.Message) {
		t.Logf("Received message: %s\n", string(msg.Value))

		var sensor entity.Sensor

		json.Unmarshal(msg.Value, &sensor)

		t.Logf("Sensor received: %#v\n", sensor)

		messagesReceived[1] = &sensor

	}

	producer.PublishMessage("prova_test_presistence", message)

	consumer1.StartConsumingMessages("prova_test_presistence", callback1, true)
	consumer2.StartConsumingMessages("prova_test_presistence", callback2, true)

	if len(messagesReceived) != 2 || messagesReceived[0] == nil || messagesReceived[1] == nil {
		t.Fatalf("Messages wasn't read")
	}

	if messagesReceived[0].IdSensor != messagesReceived[1].IdSensor || messagesReceived[0].Timestamp != messagesReceived[1].Timestamp || messagesReceived[0].TipoPoluente != messagesReceived[1].TipoPoluente || messagesReceived[0].Nivel != messagesReceived[1].Nivel {
		t.Fatal("Messages are different!")
	}

}
