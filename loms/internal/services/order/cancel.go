package order

import (
	"context"
	"route256/loms/internal/models"
)

type ReserveCanceller interface {
	ReserveCancel(ctx context.Context, items []models.ItemOrderInfo) error
}

type CancelService struct {
	orderProvide      OrderProvider
	reserveCanceller  ReserveCanceller
	orderStatusSetter OrderStatusSetter
	statusNotifier    StatusNotifier
}

type CancelDeps struct {
	OrderProvider
	ReserveCanceller
	OrderStatusSetter
	StatusNotifier
}

func NewCancelService(d CancelDeps) *CancelService {
	return &CancelService{
		orderProvide:      d.OrderProvider,
		reserveCanceller:  d.ReserveCanceller,
		orderStatusSetter: d.OrderStatusSetter,
		statusNotifier:    d.StatusNotifier,
	}
}

func (c *CancelService) CancelOrder(ctx context.Context, orderID models.OrderID) error {
	order, err := c.orderProvide.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	if err := c.reserveCanceller.ReserveCancel(ctx, order.Items); err != nil {
		return err
	}

	if err := c.statusNotifier.NotifyOrderStatus(ctx, orderID, models.StatusCancelled); err != nil {
		return err
	}

	return c.orderStatusSetter.SetStatus(ctx, orderID, models.StatusCancelled)
}
