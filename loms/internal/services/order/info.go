package order

import (
	"context"
	"route256/loms/internal/models"
)

type OrderProvider interface {
	GetOrderByID(ctx context.Context, orderID models.OrderID) (*models.Order, error)
}

type GetInfoService struct {
	orderProvider OrderProvider
}

func NewGetInfoService(orderProvider OrderProvider) *GetInfoService {
	return &GetInfoService{
		orderProvider: orderProvider,
	}
}

func (usc *GetInfoService) GetInfo(ctx context.Context, orderID models.OrderID) (*models.Order, error) {
	return usc.orderProvider.GetOrderByID(ctx, orderID)
}
