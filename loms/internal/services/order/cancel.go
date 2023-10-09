package order

import (
	"context"
	"route256/loms/internal/models"
)

type ReserveCanceller interface {
	ReserveCancel(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error
}

type CancelService struct {
	orderProvide      OrderProvider
	reserveCanceller  ReserveCanceller
	orderStatusSetter OrderStatusSetter
}

type CancelDeps struct {
	OrderProvider
	ReserveCanceller
	OrderStatusSetter
}

func NewCancelService(d CancelDeps) *CancelService {
	return &CancelService{
		orderProvide:      d.OrderProvider,
		reserveCanceller:  d.ReserveCanceller,
		orderStatusSetter: d.OrderStatusSetter,
	}
}

func (usc *CancelService) CancelOrder(ctx context.Context, orderID models.OrderID) error {
	order, err := usc.orderProvide.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	if err := usc.reserveCanceller.ReserveCancel(ctx, order.UserID); err != nil {
		return err
	}

	return usc.orderStatusSetter.SetStatus(ctx, orderID, models.StatusCancelled)
}
