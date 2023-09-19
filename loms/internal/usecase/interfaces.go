package usecase

import (
	"context"
	"route256/loms/internal/models"
)

type (
	WarehouseManagementSystem interface {
		GetStockInfo(
			ctx context.Context,
			SKU models.SKU,
		) (count uint64, err error)
	}

	OrderManagementSystem interface {
		OrderCreator
		OrderProvider
		OrderAsPaidMarker
		OrderCanceller
	}

	OrderCreator interface {
		CreateOrder(
			ctx context.Context,
			userID models.UserID,
			info CreateOrderInfo,
		) (models.OrderID, error)
	}

	OrderProvider interface {
		GetOrderInfo(
			ctx context.Context,
			orderID models.OrderID,
		) (models.Order, error)
	}

	OrderAsPaidMarker interface {
		MarkOrderAsPaid(
			ctx context.Context,
			orderID models.OrderID,
		) error
	}

	OrderCanceller interface {
		CancelOrder(
			ctx context.Context,
			orderID models.OrderID,
		) error
	}
)
