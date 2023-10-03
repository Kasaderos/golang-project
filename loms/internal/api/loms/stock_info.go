package loms

import (
	"context"
	"route256/loms/internal/models"
	servicepb "route256/loms/pkg/api/loms/v1"
)

type StockInfoService interface {
	GetStockInfo(ctx context.Context, SKU models.SKU) (count uint64, err error)
}

type StocksService interface {
	GetStockInfo(
		ctx context.Context,
		SKU models.SKU,
	) (count uint64, err error)
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
