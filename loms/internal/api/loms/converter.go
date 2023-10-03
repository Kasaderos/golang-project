package loms

import (
	"route256/loms/internal/models"
	servicepb "route256/loms/pkg/api/loms/v1"
)

func ToGetOrderInfoResponse(order *models.Order) *servicepb.GetOrderInfoResponse {
	items := make([]*servicepb.OrderInfoItem, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, &servicepb.OrderInfoItem{
			Sku:   int64(item.SKU),
			Count: uint32(item.Count),
		})
	}

	return &servicepb.GetOrderInfoResponse{
		Status: order.Status.String(),
		User:   int64(order.UserID),
		Items:  items,
	}
}
