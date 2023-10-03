package carts

import (
	"context"
	"route256/cart/internal/models"
	servicepb "route256/cart/pkg/api/carts/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ItemAddService interface {
	AddItem(ctx context.Context, userID models.UserID, sku models.SKU, count uint16) error
}

func (s Service) ItemAdd(ctx context.Context, req *servicepb.ItemAddRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	if err := s.itemAddService.AddItem(
		ctx,
		models.UserID(req.User),
		models.SKU(req.Sku),
		uint16(req.Count),
	); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
