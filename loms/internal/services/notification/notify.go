package notification

import (
	"encoding/json"
	"fmt"
	"route256/loms/internal/models"
	"route256/loms/pkg/kafka"

	"github.com/Shopify/sarama"
)

const orderStatusesTopic = "order-statuses-topic"

type message struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

type Service struct {
	producer sarama.SyncProducer
}

func NewService(producer sarama.SyncProducer) *Service {
	return &Service{
		producer: producer,
	}
}

func (c *Service) NotifyOrderStatus(orderID models.OrderID, status models.Status) error {
	m := message{
		OrderID: fmt.Sprint(orderID),
		Status:  status.String(),
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	msg, err := kafka.BuildMessage(orderStatusesTopic, m.OrderID, bytes)
	if err != nil {
		return err
	}

	_, _, err = c.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
