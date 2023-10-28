package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"route256/loms/pkg/kafka"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

var (
	topicName       = os.Getenv("ORDER_STATUSES_TOPIC")
	consumerGroupID = os.Getenv("ORDER_STATUSES_CONSUMER_GROUP_ID")
)

var brokers = []string{
	"kafka-broker-1:9091",
	"kafka-broker-2:9092",
	"kafka-broker-3:9093",
}

func main() {
	// todo
	// autocreate kafka cluster
	// заходим в кафка ui руками заполняем конфиг на все 2 мин
	if os.Getenv("WAIT_KAFKA") == "on" {
		time.Sleep(time.Minute * 2)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Наш обработчик реализующий интерфейс sarama.ConsumerGroupHandler
	consumerGroupHandler := NewConsumerGroupHandler()

	// Создаем коньюмер группу
	consumerGroup, err := kafka.NewConsumerGroup(
		brokers,
		consumerGroupID,
		[]string{topicName},
		consumerGroupHandler,
	)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		// запускаем вычитку сообщений
		consumerGroup.Run(ctx)
	}()

	<-consumerGroupHandler.Ready() // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	var (
		consumptionIsPaused = false
		keepRunning         = true
	)
	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(consumerGroup, &consumptionIsPaused)
		}
	}

	cancel()
	wg.Wait()

	if err = consumerGroup.Close(); err != nil {
		log.Fatalf("Error closing consumer group: %v", err)
	}
}

func toggleConsumptionFlow(cg sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		cg.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		cg.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}
