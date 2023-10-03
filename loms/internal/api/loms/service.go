package loms

import (
	servicepb "route256/loms/pkg/api/loms/v1"
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

func NewServer(d *Deps) *Service {
	return &Service{
		orderCreateService: d.OrderCreateService,
		orderInfoService:   d.OrderInfoService,
		orderPayService:    d.OrderPayService,
		orderCancelService: d.OrderCancelService,
		stockInfoService:   d.StockInfoService,
	}
}
