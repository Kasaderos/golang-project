package services

import (
	"context"
	"route256/cart/internal/models"
)

type CartService interface {
	AddItem(ctx context.Context, userID models.UserID, sku models.SKU, count uint16) error
	ListItem(
		ctx context.Context,
		userID models.UserID,
	) (totalPrice uint32, items []models.CartItem, err error)
	Checkout(ctx context.Context, userID models.UserID) (models.OrderID, error)
	DeleteItem(ctx context.Context, userID models.UserID, sku models.SKU) error
	Clear(ctx context.Context, userID models.UserID) error
}
