package mock

import (
	"context"
	"route256/cart/internal/models"
)

type CartRepository struct {
}

func NewCartRepostiory() *CartRepository {
	return &CartRepository{}
}

func (r *CartRepository) AddItem(
	ctx context.Context,
	userID models.UserID,
	item models.CartItem,
) error {
	return nil
}

func (r *CartRepository) GetItemsByUserID(
	ctx context.Context,
	userID models.UserID,
) ([]models.CartItem, error) {
	return []models.CartItem{}, nil
}

func (r *CartRepository) DeleteItem(
	ctx context.Context,
	userID models.UserID,
	SKU models.SKU,
) error {
	return nil
}

func (r *CartRepository) DeleteItemsByUserID(
	ctx context.Context,
	userID models.UserID,
) error {
	return nil
}
