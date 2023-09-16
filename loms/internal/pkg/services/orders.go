package services

import (
	"context"
	"log"
	"route256/loms/internal/pkg/models"
)

const OrderStatusFailed = "failed"
const OrderStatusAwaiting = "awaiting payment"

type OrdersStorage interface {
	Create(ctx context.Context, order models.Order) (orderID string, err error)
	SetStatus(ctx context.Context, status string) (err error)
}

type StocksReserver interface {
	Reserve(ctx context.Context, order models.Order) error
}

type OrderService struct {
	orderStorage   OrdersStorage
	stocksReserver StocksReserver
}

func (o *OrderService) Create(ctx context.Context, order models.Order) (orderID string, err error) {
	orderID, err = o.orderStorage.Create(ctx, order)
	if err != nil {
		return "", err
	}

	var (
		status     string
		errReserve error
	)

	errReserve = o.stocksReserver.Reserve(ctx, order)
	if errReserve != nil {
		status = OrderStatusFailed
	} else {
		status = OrderStatusAwaiting
	}

	if err := o.orderStorage.SetStatus(ctx, status); err != nil {
		log.Println("while setting status", err)
		return "", errReserve
	}

	return orderID, errReserve
}

func (o *OrderService) Info()   {}
func (o *OrderService) Pay()    {}
func (o *OrderService) Cancel() {}
