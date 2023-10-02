package server

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
