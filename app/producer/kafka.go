package producer

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewProducer(server string) *kafka.Producer {
	producer, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": server,
			"security.protocol": "PLAINTEXT",
		},
	)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	return producer
}
