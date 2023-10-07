package stock

import (
	"context"
	"route256/loms/internal/models"
)

type StockProvider interface {
	GetStocksBySKU(ctx context.Context, SKU models.SKU) (count uint64, err error)
}

type StocksService struct {
	StockProvider
}

func NewStocksService(provider StockProvider) *StocksService {
	return &StocksService{
		StockProvider: provider,
	}
}

func (usc *StocksService) GetStockInfo(
	ctx context.Context,
	SKU models.SKU,
) (count uint64, err error) {
	return usc.StockProvider.GetStocksBySKU(ctx, SKU)
}
