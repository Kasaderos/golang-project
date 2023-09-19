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
		StocksReserver
		ReserveRemover
		ReserveCanceller
	}

	StocksReserver interface {
		ReserveStocks(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error
	}

	ReserveRemover interface {
		ReserveRemove(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error
	}

	ReserveCanceller interface {
		ReserveCancel(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error
	}

	// order management system repository
	OMSRepository interface {
		OrderCreator
		OrderStatusSetter
		OrderProvider
	}

	OrderCreator interface {
		CreateOrder(ctx context.Context, order models.Order) (models.Order, error)
	}

	OrderStatusSetter interface {
		SetStatus(ctx context.Context, status models.Status) error
	}

	OrderProvider interface {
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

func (usc *omsUsecase) CreateOrder(
	ctx context.Context,
	userID models.UserID,
	info usecase.CreateOrderInfo,
) (models.OrderID, error) {
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
		return models.OrderID(-1), err
	}

	err = usc.WMSRepository.ReserveStocks(ctx, userID, info.Items)
	if err != nil {
		return models.OrderID(-1), err
	}

	return OrderID, nil
}

func (usc *omsUsecase) GetOrderInfo(
	ctx context.Context,
	orderID models.OrderID,
) (models.Order, error) {
	return usc.OMSRepository.GetOrderByID(ctx, orderID)
}

func (usc *omsUsecase) MarkOrderAsPaid(
	ctx context.Context,
	orderID models.OrderID,
) error {
	order, err := usc.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	if err := usc.WMSRepository.ReserveRemove(ctx, order.UserID, order.Items); err != nil {
		return err
	}

	return usc.OMSRepository.SetStatus(ctx, models.StatusPaid)
}

func (usc *omsUsecase) CancelOrder(
	ctx context.Context,
	orderID models.OrderID,
) error {
	order, err := usc.OMSRepository.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	if err := usc.WMSRepository.ReserveCancel(ctx, order.UserID, order.Items); err != nil {
		return err
	}

	return usc.OMSRepository.SetStatus(ctx, models.StatusCancelled)
}
