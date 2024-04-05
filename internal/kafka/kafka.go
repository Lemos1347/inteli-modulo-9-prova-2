package kafka

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumerCallBack func(msg *kafka.Message)

type KafkaClientType uint8

const (
	Producer KafkaClientType = iota
	Consumer
)

func (s KafkaClientType) String() string {
	switch s {
	case Producer:
		return "Producer"
	case Consumer:
		return "Consumer"
	}
	return "unknown"
}

type KafkaClient struct {
	Producer *kafka.Producer
	Consumer *kafka.Consumer
}

func (s *KafkaClient) PublishMessage(topic string, message []byte) {
	if s.Producer == nil {
		log.Panic("Trying to publish message but producer doesn't exists")
	}

	s.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	log.Printf("Message produced: %#v \n", message)

	s.Producer.Flush(30 * 1000)
}

func (s *KafkaClient) StartConsumingMessages(topic string, callback ConsumerCallBack, justOne ...any) {
	if s.Consumer == nil {
		log.Panic("Trying to consume message but consumer doesn't exists")
	}

	s.Consumer.SubscribeTopics([]string{topic}, nil)

	log.Printf("Consuming messages from: %s\n", topic)

	defer s.Consumer.Close()

	for {
		msg, err := s.Consumer.ReadMessage(-1)
		if err == nil {
			callback(msg)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
		if len(justOne) > 0 {
			log.Print("Configurado para consumir so uma vez")
			break
		}
	}
}

func generateProducer() *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"client.id":         "go-producer",
	})
	if err != nil {
		log.Panicf("Error while trying to create producer: %s\n", err.Error())
	}

	return producer
}

func generateConsumer() *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Panicf("Error while trying to create consumer: %s\n", err.Error())
	}

	return consumer
}

func generatePersonalizedConsumer(configMap *kafka.ConfigMap) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Panicf("Error while trying to create consumer: %s\n", err.Error())
	}

	return consumer
}

func GenerateKafkaClient(clientType KafkaClientType, configMap ...*kafka.ConfigMap) *KafkaClient {
	log.Printf("Generating kafka client for: %s\n", clientType)

	kafkaClient := KafkaClient{}

	switch clientType {
	case Producer:
		kafkaClient.Producer = generateProducer()

	case Consumer:
		if len(configMap) > 0 {
			if configMap[0] == nil {
				log.Fatal("Trying to create personalized kafka, but config is nil")
			}
			kafkaClient.Consumer = generatePersonalizedConsumer(configMap[0])

		} else {
			kafkaClient.Consumer = generateConsumer()
		}

	default:
		log.Panic("Undefined kafka client type")
	}

	if kafkaClient.Producer == nil && kafkaClient.Consumer == nil {
		log.Panic("Neither consumer or producer was created!")
	}

	log.Printf("%s created!\n", clientType)

	return &kafkaClient
}
