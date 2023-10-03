package loms

import (
	"context"
	"route256/loms/internal/models"
	servicepb "route256/loms/pkg/api/loms/v1"
)

type OrderCreateService interface {
	CreateOrder(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) (models.OrderID, error)
}

func (s Service) OrderCreate(ctx context.Context, req *servicepb.OrderCreateRequest) (*servicepb.OrderCreateResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	userID, items := FromOrderCreateRequest(req)

	orderID, err := s.orderCreateService.CreateOrder(
		ctx,
		userID,
		items,
	)
	if err != nil {
		return nil, err
	}

	return &servicepb.OrderCreateResponse{
		OrderId: int64(orderID),
	}, nil
}

func FromOrderCreateRequest(req *servicepb.OrderCreateRequest) (models.UserID, []models.ItemOrderInfo) {
	items := make([]models.ItemOrderInfo, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, models.ItemOrderInfo{
			SKU:   models.SKU(item.Sku),
			Count: uint16(item.Count),
		})
	}
	return models.UserID(req.User), items
}
