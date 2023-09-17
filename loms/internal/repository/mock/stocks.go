package repository

import (
	"context"
	"route256/loms/internal/models"
	oms "route256/loms/internal/usecase/OMS"
)

type stocksRepository struct {
}

var _ oms.WMSRepository = (*stocksRepository)(nil)

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
	return 0, nil
}
