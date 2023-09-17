package usecase

import (
	"context"
	"route256/loms/internal/models"
)

type OrderManagementSystem interface {
	CreateOrder(
		ctx context.Context,
		userID models.UserID,
		info CreateOrderInfo,
	) (models.OrderID, error)
	GetOrderInfo(
		ctx context.Context,
		orderID models.OrderID,
	) (models.Order, error)
	MarkOrderAsPaid(
		ctx context.Context,
		orderID models.OrderID,
	) error
	CancelOrder(
		ctx context.Context,
		orderID models.OrderID,
	) error
}
