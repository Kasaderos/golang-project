package order

import (
	"context"
	"route256/loms/internal/models"
)

type OrderStatusSetter interface {
	SetStatus(ctx context.Context, orderID models.OrderID, status models.Status) error
}

type ReserveRemover interface {
	ReserveRemove(ctx context.Context, userID models.UserID) error
}

type PayService struct {
	orderProvider     OrderProvider
	reserveRemover    ReserveRemover
	orderStatusSetter OrderStatusSetter
}

type PayDeps struct {
	OrderProvider
	ReserveRemover
	OrderStatusSetter
}

func NewPayService(d PayDeps) *PayService {
	return &PayService{
		orderProvider:     d.OrderProvider,
		reserveRemover:    d.ReserveRemover,
		orderStatusSetter: d.OrderStatusSetter,
	}
}

func (usc *PayService) MarkAsPaid(
	ctx context.Context,
	orderID models.OrderID,
) error {
	order, err := usc.orderProvider.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	if err := usc.reserveRemover.ReserveRemove(ctx, order.UserID); err != nil {
		return err
	}

	return usc.orderStatusSetter.SetStatus(ctx, orderID, models.StatusPaid)
}
