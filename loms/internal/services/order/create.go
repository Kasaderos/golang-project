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
	orderCreator   OrderCreator
	stocksReserver StocksReserver
}

type CreateDeps struct {
	OrderCreator
	StocksReserver
}

func NewCreateService(d CreateDeps) *CreateService {
	return &CreateService{
		orderCreator:   d.OrderCreator,
		stocksReserver: d.StocksReserver,
	}
}

func (usc *CreateService) CreateOrder(
	ctx context.Context,
	userID models.UserID,
	items []models.ItemOrderInfo,
) (models.OrderID, error) {
	order := models.Order{
		UserID: userID,
		Items:  items,
	}

	orderID, err := usc.orderCreator.CreateOrder(ctx, order)
	if err != nil {
		return models.OrderID(-1), err
	}

	if err := usc.stocksReserver.ReserveStocks(ctx, items); err != nil {
		return models.OrderID(-1), err
	}

	return orderID, nil
}
