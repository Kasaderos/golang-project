package app

import (
	"fmt"
	"route256/loms/pkg/kafka"
	"time"

	"github.com/Shopify/sarama"
)

var brokers = []string{
	"127.0.0.1:9091",
	"127.0.0.1:9092",
	"127.0.0.1:9093",
}

const (
	MaxOpenRequests = 1
	MaxRetries      = 5
	RetryBackoff    = 10 * time.Millisecond
)

func initKafkaProducer() (sarama.SyncProducer, error) {
	producer, err := kafka.NewSyncProducer(brokers,
		kafka.WithIdempotent(),
		kafka.WithRequiredAcks(sarama.WaitForLocal),
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
