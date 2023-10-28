package order

//go:generate mkdir -p mock
//go:generate minimock -o ./mock/ -s .go -g

import (
	"context"
	"log"
	"route256/loms/internal/models"
	"time"
)

const NotifyTimeout = time.Minute * 2

type OrderCreator interface {
	CreateOrder(ctx context.Context, order models.Order) (models.OrderID, error)
}

type StocksReserver interface {
	ReserveStocks(ctx context.Context, items []models.ItemOrderInfo) error
}

type StatusNotifier interface {
	NotifyOrderStatus(order models.OrderID, status models.Status) error
}

type CreateService struct {
	orderCreator      OrderCreator
	stocksReserver    StocksReserver
	orderStatusSetter OrderStatusSetter
	statusNotifier    StatusNotifier
}

type CreateDeps struct {
	OrderCreator
	StocksReserver
	OrderStatusSetter
	StatusNotifier
}

func NewCreateService(d CreateDeps) *CreateService {
	return &CreateService{
		orderCreator:      d.OrderCreator,
		stocksReserver:    d.StocksReserver,
		orderStatusSetter: d.OrderStatusSetter,
	}
}

func (c *CreateService) CreateOrder(
	ctx context.Context,
	userID models.UserID,
	items []models.ItemOrderInfo,
) (models.OrderID, error) {
	order := models.Order{
		UserID: userID,
		Status: models.StatusNew,
		Items:  items,
	}

	orderID, err := c.orderCreator.CreateOrder(ctx, order)
	if err != nil {
		return models.OrderID(-1), err
	}

	go func() {
		if err := c.statusNotifier.NotifyOrderStatus(orderID, models.StatusNew); err != nil {
			log.Println("notifier: %w", err)
			// save somehow and then somehow notify
		}
	}()

	if err := c.stocksReserver.ReserveStocks(ctx, items); err != nil {
		if err := c.orderStatusSetter.SetStatus(
			ctx,
			orderID,
			models.StatusFailed,
		); err != nil {
			return models.OrderID(-1), err
		}
		return models.OrderID(-1), err
	}

	if err := c.orderStatusSetter.SetStatus(
		ctx,
		orderID,
		models.StatusAwaitingPayment,
	); err != nil {
		return models.OrderID(-1), err
	}

	return orderID, nil
}
