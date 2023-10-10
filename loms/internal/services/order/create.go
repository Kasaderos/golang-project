package order

import (
	"context"
	"route256/loms/internal/models"
)

type OrderCreator interface {
	CreateOrder(ctx context.Context, order models.Order) (models.OrderID, error)
}

type StocksReserver interface {
	ReserveStocks(ctx context.Context, items []models.ItemOrderInfo) error
}

type CreateService struct {
	orderCreator      OrderCreator
	stocksReserver    StocksReserver
	orderStatusSetter OrderStatusSetter
}

type CreateDeps struct {
	OrderCreator
	StocksReserver
	OrderStatusSetter
}

func NewCreateService(d CreateDeps) *CreateService {
	return &CreateService{
		orderCreator:      d.OrderCreator,
		stocksReserver:    d.StocksReserver,
		orderStatusSetter: d.OrderStatusSetter,
	}
}

func (usc *CreateService) CreateOrder(
	ctx context.Context,
	userID models.UserID,
	items []models.ItemOrderInfo,
) (models.OrderID, error) {
	order := models.Order{
		UserID: userID,
		Status: models.StatusNew,
		Items:  items,
	}

	orderID, err := usc.orderCreator.CreateOrder(ctx, order)
	if err != nil {
		return models.OrderID(-1), err
	}

	if err := usc.stocksReserver.ReserveStocks(ctx, items); err != nil {
		if err := usc.orderStatusSetter.SetStatus(
			ctx,
			orderID,
			models.StatusFailed,
		); err != nil {
			return models.OrderID(-1), err
		}
		return models.OrderID(-1), err
	}

	if err := usc.orderStatusSetter.SetStatus(
		ctx,
		orderID,
		models.StatusAwaitingPayment,
	); err != nil {
		return models.OrderID(-1), err
	}

	return orderID, nil
}
