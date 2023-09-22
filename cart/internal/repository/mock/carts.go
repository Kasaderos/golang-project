package mock

import (
	"context"
	"route256/cart/internal/models"
)

type cartRepository struct {
}

// var _ oms.OMSRepository = (*cartRepository)(nil)

func NewCartRepostiory() *cartRepository {
	return &cartRepository{}
}

func (r *cartRepository) AddItem(
	ctx context.Context,
	userID models.UserID,
	item models.CartItem,
) error {
	return nil
}

func (r *cartRepository) GetItemsByUserID(
	ctx context.Context,
	userID models.UserID,
) ([]models.CartItem, error) {
	return []models.CartItem{}, nil
}

func (r *cartRepository) DeleteItem(
	ctx context.Context,
	userID models.UserID,
	SKU models.SKU,
) error {
	return nil
}

func (r *cartRepository) DeleteItemsByUserID(
	ctx context.Context,
	userID models.UserID,
) error {
	return nil
}
