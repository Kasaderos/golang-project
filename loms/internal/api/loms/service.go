package loms

import (
	"context"
	"route256/loms/internal/models"
	usecase "route256/loms/internal/services"
	"route256/loms/internal/services/order"
	"route256/loms/internal/services/stock"
	"route256/loms/pkg/api/loms/v1"
	servicepb "route256/loms/pkg/api/loms/v1"
)

var _ servicepb.LOMSServer = (*Service)(nil)

// Service - уровень Delivery
type Service struct {
	servicepb.UnimplementedLOMSServer

	orderCreateService *order.CreateService
	orderInfoService   *order.GetInfoService
	orderPayService    *order.PayService
	orderCancelService *order.CancelService
	stockInfoService   *stock.StocksService
}

func (s Service) CreateOrder(ctx context.Context, req *loms.CreateOrderRequest) (*loms.CreateOrderResponse, error) {
	items := make([]models.ItemOrderInfo, 0, len(req.Order.Items))
	for _, item := range req.Order.Items {
		items = append(items, models.ItemOrderInfo{
			SKU:   models.SKU(item.Sku),
			Count: uint16(item.Count),
		})
	}

	orderID, err := s.orderCreateService.CreateOrder(
		ctx,
		models.UserID(req.Order.UserId),
		usecase.CreateOrderInfo{
			Items: items,
		},
	)
	if err != nil {
		return nil, err
	}

	return &servicepb.CreateOrderResponse{
		OrderId: int64(orderID),
	}, nil
}

func (s Service) GetStockInfo(ctx context.Context, req *loms.GetStockInfoRequest) (*loms.GetStockInfoResponse, error) {
	count, err := s.stockInfoService.GetStockInfo(ctx, models.SKU(req.Sku))
	if err != nil {
		return nil, err
	}

	return &servicepb.GetStockInfoResponse{
		Count: count,
	}, nil
}
