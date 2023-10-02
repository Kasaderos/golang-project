package server

import (
	"route256/cart/internal/models"
	servicepb "route256/cart/pkg/api/carts/v1"
)

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
