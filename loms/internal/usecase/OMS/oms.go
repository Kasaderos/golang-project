package oms

import (
	"context"
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
		ExpiredOrdersProvider
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

	ExpiredOrdersProvider interface {
		ListExpiredOrders(ctx context.Context, limit uint32) ([]models.Order, error)
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
