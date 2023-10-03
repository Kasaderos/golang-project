package carts

import (
	"context"
	"route256/cart/internal/models"
	servicepb "route256/cart/pkg/api/carts/v1"
)

type ListItemService interface {
	ListItem(ctx context.Context, userID models.UserID) (totalPrice uint32, items []models.CartItem, err error)
}

func (s Service) List(ctx context.Context, req *servicepb.ListRequest) (*servicepb.ListResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	totalPrice, items, err := s.listItemService.ListItem(
		ctx,
		models.UserID(req.User),
	)
	if err != nil {
		return nil, err
	}

	return ToListResponse(totalPrice, items), nil
}

func ToListResponse(totalPrice uint32, cartItems []models.CartItem) *servicepb.ListResponse {
	items := make([]*servicepb.ListItem, 0, len(cartItems))
	for _, item := range cartItems {
		items = append(items, &servicepb.ListItem{
			Sku:   uint32(item.SKU),
			Count: uint32(item.Count),
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return &servicepb.ListResponse{
		TotalPrice: totalPrice,
		Items:      items,
	}
}
