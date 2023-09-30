package loms

import (
	"context"
	"route256/loms/internal/models"
	dto "route256/loms/internal/services"
	servicepb "route256/loms/pkg/api/loms/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

var _ servicepb.LOMSServer = (*Service)(nil)

type Service struct {
	servicepb.UnimplementedLOMSServer

	orderCreateService OrderCreateService
	orderInfoService   OrderInfoService
	orderPayService    OrderPayService
	orderCancelService OrderCancelService
	stockInfoService   StockInfoService
}

type Deps struct {
	OrderCreateService
	OrderInfoService
	OrderPayService
	OrderCancelService
	StockInfoService
}

func NewServer(d Deps) *Service {
	return &Service{
		orderCreateService: d.OrderCreateService,
		orderInfoService:   d.OrderInfoService,
		orderPayService:    d.OrderPayService,
		orderCancelService: d.OrderCancelService,
		stockInfoService:   d.StockInfoService,
	}
}

type (
	StocksService interface {
		GetStockInfo(
			ctx context.Context,
			SKU models.SKU,
		) (count uint64, err error)
	}

	OrderCancelService interface {
		CancelOrder(
			ctx context.Context,
			orderID models.OrderID,
		) error
	}

	OrderCreateService interface {
		CreateOrder(ctx context.Context, userID models.UserID, info dto.CreateOrderInfo) (models.OrderID, error)
	}

	OrderPayService interface {
		MarkAsPaid(ctx context.Context, orderID models.OrderID) error
	}

	OrderInfoService interface {
		GetInfo(ctx context.Context, orderID models.OrderID) (*models.Order, error)
	}

	StockInfoService interface {
		GetStockInfo(ctx context.Context, SKU models.SKU) (count uint64, err error)
	}
)

func (s Service) OrderCreate(ctx context.Context, req *servicepb.OrderCreateRequest) (*servicepb.OrderCreateResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	items := make([]models.ItemOrderInfo, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, models.ItemOrderInfo{
			SKU:   models.SKU(item.Sku),
			Count: uint16(item.Count),
		})
	}

	orderID, err := s.orderCreateService.CreateOrder(
		ctx,
		models.UserID(req.UserId),
		dto.CreateOrderInfo{
			Items: items,
		},
	)
	if err != nil {
		return nil, err
	}

	return &servicepb.OrderCreateResponse{
		OrderId: int64(orderID),
	}, nil
}

func (s Service) GetStockInfo(ctx context.Context, req *servicepb.GetStockInfoRequest) (*servicepb.GetStockInfoResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	count, err := s.stockInfoService.GetStockInfo(ctx, models.SKU(req.Sku))
	if err != nil {
		return nil, err
	}

	return &servicepb.GetStockInfoResponse{
		Count: count,
	}, nil
}

func (c *Service) CancelOrder(ctx context.Context, req *servicepb.CancelOrderRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	if err := c.orderCancelService.CancelOrder(
		ctx,
		models.OrderID(req.OrderId),
	); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (c *Service) GetOrderInfo(ctx context.Context, req *servicepb.GetOrderInfoRequest) (*servicepb.GetOrderInfoResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	order, err := c.orderInfoService.GetInfo(
		ctx,
		models.OrderID(req.OrderId),
	)
	if err != nil {
		return nil, err
	}

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
	}, nil
}

func (c *Service) OrderPay(ctx context.Context, req *servicepb.OrderPayRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	if err := c.orderPayService.MarkAsPaid(
		ctx,
		models.OrderID(req.OrderId),
	); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
