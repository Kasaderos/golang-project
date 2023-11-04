package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"route256/loms/internal/models"
	"route256/loms/pkg/kafka"

	"github.com/Shopify/sarama"
)

const orderStatusesTopic = "order-statuses-topic"

const messageQueueSize = 10

type message struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
	bytes   []byte
	ctx     context.Context
}

type Service struct {
	producer sarama.SyncProducer
	messages chan message
}

func NewService(producer sarama.SyncProducer) *Service {
	return &Service{
		producer: producer,
		messages: make(chan message, messageQueueSize),
	}
}

func (c *Service) NotifyOrderStatus(ctx context.Context, orderID models.OrderID, status models.Status) error {
	m := message{
		OrderID: fmt.Sprint(orderID),
		Status:  status.String(),
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	m.ctx = ctx
	m.bytes = bytes

	select {
	case <-ctx.Done():
		return ctx.Err()
	case c.messages <- m:
		// ok queued
	}

	return nil
}

func (c *Service) NotifyOrderStatusBackground(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case m := <-c.messages:
			msg, err := kafka.BuildMessage(orderStatusesTopic, m.OrderID, m.bytes)
			if err != nil {
				log.Println("notifications: build message:", err)
				c.retrySend(m)
				continue
			}

			_, _, err = c.producer.SendMessage(msg)
			if err != nil {
				log.Println("notifications: send message:", err)
				c.retrySend(m)
			}
		}
	}
}

func (c *Service) retrySend(m message) {
	select {
	case c.messages <- m:
		// ok queued
	case <-m.ctx.Done():
		// save somehow
	}
}
