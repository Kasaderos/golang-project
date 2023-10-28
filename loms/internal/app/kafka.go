package app

import (
	"fmt"
	"route256/loms/pkg/kafka"
	"time"

	"github.com/Shopify/sarama"
)

var brokers = []string{
	"kafka-broker-1:9091",
	"kafka-broker-2:9092",
	"kafka-broker-3:9093",
}

const (
	MaxOpenRequests = 1
	MaxRetries      = 5
	RetryBackoff    = 10 * time.Millisecond
)

func initKafkaProducer() (sarama.SyncProducer, error) {
	producer, err := kafka.NewSyncProducer(brokers,
		kafka.WithIdempotent(),
		kafka.WithRequiredAcks(sarama.WaitForAll),
		kafka.WithProducerPartitioner(sarama.NewRoundRobinPartitioner),
		kafka.WithMaxOpenRequests(MaxOpenRequests),
		kafka.WithMaxRetries(MaxRetries),
		kafka.WithRetryBackoff(RetryBackoff),
	)
	if err != nil {
		return nil, fmt.Errorf("kafka: %w", err)
	}

	return producer, nil
}
