package loms

import (
	"context"
	"route256/loms/internal/models"
	servicepb "route256/loms/pkg/api/loms/v1"
)

type OrderInfoService interface {
	GetInfo(ctx context.Context, orderID models.OrderID) (*models.Order, error)
}

func (c *Service) GetOrderInfo(ctx context.Context, req *servicepb.GetOrderInfoRequest) (*servicepb.GetOrderInfoResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	order, err := c.orderInfoService.GetInfo(
		ctx,
		models.OrderID(req.OrderId),
	)
	if err != nil {
		return nil, err
	}

	return ToGetOrderInfoResponse(order), nil
}
