package carts

import (
	"context"
	"route256/cart/internal/models"
	servicepb "route256/cart/pkg/api/carts/v1"
)

type CheckoutService interface {
	Checkout(ctx context.Context, userID models.UserID) (models.OrderID, error)
}

func (s Service) Checkout(ctx context.Context, req *servicepb.CheckoutRequest) (*servicepb.CheckoutResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	orderID, err := s.checkoutService.Checkout(
		ctx,
		models.UserID(req.User),
	)
	if err != nil {
		return nil, err
	}

	return &servicepb.CheckoutResponse{
		OrderId: int64(orderID),
	}, nil
}
