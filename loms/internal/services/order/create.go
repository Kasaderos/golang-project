package order

import (
	"context"
	"math/rand"
	"route256/loms/internal/models"
	dto "route256/loms/internal/services"
)

type OrderCreator interface {
	CreateOrder(ctx context.Context, order models.Order) error
}

type StocksReserver interface {
	ReserveStocks(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error
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
	info dto.CreateOrderInfo,
) (models.OrderID, error) {
	var (
		OrderID = models.OrderID(rand.Int() % 1000)
		order   = models.Order{
			ID:     OrderID,
			UserID: userID,
			Items:  info.Items,
		}
	)

	if err := usc.orderCreator.CreateOrder(ctx, order); err != nil {
		return models.OrderID(-1), err
	}

	if err := usc.stocksReserver.ReserveStocks(ctx, userID, info.Items); err != nil {
		return models.OrderID(-1), err
	}

	return OrderID, nil
}
