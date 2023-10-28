package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

func NewAsyncProducer(brokers []string, opts ...Option) (sarama.AsyncProducer, error) {
	config, err := prepareProducerSaramaConfig(opts...)
	if err != nil {
		return nil, err
	}

	asyncProducer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, errors.Wrap(err, "error with async kafka-producer")
	}

	// !!!ВНИМАНИЕ!!!
	// ОБЯЗАТЕЛЬНОЕ чтение канала ошибок при c.Producer.Return.Errors = true
	go func() {
		// Error и Retry топики можно использовать при получении ошибки
		for err := range asyncProducer.Errors() {
			fmt.Println(err.Error())
		}
	}()

	// !!!ВНИМАНИЕ!!!
	// ОБЯЗАТЕЛЬНОЕ чтение канала успешных событий при c.Producer.Return.Successes = true
	go func() {
		for msg := range asyncProducer.Successes() {
			fmt.Println("Async success with key", msg.Key)
		}
	}()

	return asyncProducer, nil
}
