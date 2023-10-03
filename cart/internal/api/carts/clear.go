package carts

import (
	"context"
	"route256/cart/internal/models"
	servicepb "route256/cart/pkg/api/carts/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ClearService interface {
	Clear(ctx context.Context, userID models.UserID) error
}

func (s Service) Clear(ctx context.Context, req *servicepb.ClearRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	if err := s.clearService.Clear(
		ctx,
		models.UserID(req.User),
	); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
