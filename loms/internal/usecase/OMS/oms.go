package oms

import (
	"context"
	"math/rand"
	"route256/loms/internal/models"
	"route256/loms/internal/usecase"
)

type (
	// warehouse management system repository
	WMSRepository interface {
		ReserveStocks(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error
	}

	// order management system repository
	OMSRepository interface {
		CreateOrder(ctx context.Context, order models.Order) (models.Order, error)
		GetOrderByID(ctx context.Context, orderID models.OrderID) (models.Order, error)
	}
)

type Deps struct {
	WMSRepository
	OMSRepository
}

type omsUsecase struct {
	Deps
}

var _ usecase.OrderManagementSystem = (*omsUsecase)(nil)

func NewOMSUsecase(d Deps) *omsUsecase {
	return &omsUsecase{
		Deps: d,
	}
}

func (usc *omsUsecase) CreateOrder(ctx context.Context, userID models.UserID, info usecase.CreateOrderInfo) (models.OrderID, error) {
	var (
		OrderID = models.OrderID(rand.Int() % 1000)
		order   = models.Order{
			ID:     OrderID,
			UserID: userID,
			Items:  info.Items,
		}
	)

	_, err := usc.OMSRepository.CreateOrder(ctx, order)
	if err != nil {
		return models.OrderID(-1), usecase.ErrCreateOrder
	}

	err = usc.WMSRepository.ReserveStocks(ctx, userID, info.Items)
	if err != nil {
		return models.OrderID(-1), usecase.ErrReserveStocks
	}

	return OrderID, nil
}

func (usc *omsUsecase) OrderInfo(ctx context.Context, orderID models.OrderID) (models.Order, error) {
	return usc.OMSRepository.GetOrderByID(ctx, orderID)
}
