package loms

import (
	"route256/cart/internal/models"
	loms_v1 "route256/loms/pkg/api/loms/v1"
)

func ToOrderCreateRequest(userID models.UserID, items []models.CartItem) *loms_v1.OrderCreateRequest {
	reqItems := make([]*loms_v1.OrderInfoItem, 0, len(items))
	for _, item := range items {
		reqItems = append(reqItems, &loms_v1.OrderInfoItem{
			Sku:   int64(item.SKU),
			Count: uint32(item.Count),
		})
	}
	return &loms_v1.OrderCreateRequest{
		User:  int64(userID),
		Items: reqItems,
	}
}

func ToGetStockRequest(sku models.SKU) *loms_v1.GetStockInfoRequest {
	return &loms_v1.GetStockInfoRequest{
		Sku: uint32(sku),
	}
}
