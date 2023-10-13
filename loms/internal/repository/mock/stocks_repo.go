package repository

import (
	"context"
	"route256/loms/internal/models"
)

type StocksRepository struct {
}

func NewStocksRepostiory() *StocksRepository {
	return &StocksRepository{}
}

func (r *StocksRepository) ReserveStocks(
	ctx context.Context,
	userID models.UserID,
	items []models.ItemOrderInfo,
) error {
	return nil
}

func (r *StocksRepository) ReserveRemove(
	ctx context.Context,
	userID models.UserID,
) error {
	return nil
}

func (r *StocksRepository) ReserveCancel(
	ctx context.Context,
	userID models.UserID,
	items []models.ItemOrderInfo,
) error {
	return nil
}

func (r *StocksRepository) GetStocksBySKU(
	ctx context.Context,
	SKU models.SKU,
) (count uint64, err error) {
	return 100, nil
}
