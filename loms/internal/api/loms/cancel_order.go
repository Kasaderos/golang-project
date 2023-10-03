package loms

import (
	"context"
	"route256/loms/internal/models"
	servicepb "route256/loms/pkg/api/loms/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderCancelService interface {
	CancelOrder(
		ctx context.Context,
		orderID models.OrderID,
	) error
}

func (c *Service) CancelOrder(ctx context.Context, req *servicepb.CancelOrderRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	if err := c.orderCancelService.CancelOrder(
		ctx,
		models.OrderID(req.OrderId),
	); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
