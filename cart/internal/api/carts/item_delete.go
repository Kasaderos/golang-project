package carts

import (
	"context"
	"route256/cart/internal/models"

	servicepb "route256/cart/pkg/api/carts/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ItemDeleteService interface {
	DeleteItem(ctx context.Context, userID models.UserID, sku models.SKU) error
}

func (s Service) ItemDelete(ctx context.Context, req *servicepb.ItemDeleteRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	if err := s.itemDeleteService.DeleteItem(
		ctx,
		models.UserID(req.User),
		models.SKU(req.Sku),
	); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
