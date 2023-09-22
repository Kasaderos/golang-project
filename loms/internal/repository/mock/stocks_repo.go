package repository

import (
	"context"
	"route256/loms/internal/models"
)

type stocksRepository struct {
}

func NewStocksRepostiory() *stocksRepository {
	return &stocksRepository{}
}

func (r *stocksRepository) ReserveStocks(
	ctx context.Context,
	userID models.UserID,
	items []models.ItemOrderInfo,
) error {
	return nil
}

func (r *stocksRepository) ReserveRemove(
	ctx context.Context,
	userID models.UserID,
	items []models.ItemOrderInfo,
) error {
	return nil
}

func (r *stocksRepository) ReserveCancel(
	ctx context.Context,
	userID models.UserID,
	items []models.ItemOrderInfo,
) error {
	return nil
}

func (r *stocksRepository) GetStockBySKU(
	ctx context.Context,
	SKU models.SKU,
) (count uint64, err error) {
	return 100, nil
}
