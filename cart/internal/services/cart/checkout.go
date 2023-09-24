package cart

import (
	"context"
	"fmt"
	"route256/cart/internal/models"
)

type OrderCreator interface {
	CreateOrder(ctx context.Context, userID models.UserID, items []models.CartItem) (models.OrderID, error)
}

type CheckoutService struct {
	orderCreator  OrderCreator
	itemsProvider ItemsProvider
	itemsDeleter  ItemsDeleter
}

type CheckoutDeps struct {
	OrderCreator
	ItemsProvider
	ItemsDeleter
}

func NewCheckoutService(d CheckoutDeps) *CheckoutService {
	return &CheckoutService{
		orderCreator:  d.OrderCreator,
		itemsProvider: d.ItemsProvider,
		itemsDeleter:  d.ItemsDeleter,
	}
}

func (c CheckoutService) Checkout(ctx context.Context, userID models.UserID) (models.OrderID, error) {
	items, err := c.itemsProvider.GetItemsByUserID(ctx, userID)
	if err != nil {
		return models.OrderID(0), fmt.Errorf("get items: %w", err)
	}

	orderID, err := c.orderCreator.CreateOrder(ctx, userID, items)
	if err != nil {
		return models.OrderID(0), fmt.Errorf("create order: %w", err)
	}

	if err := c.itemsDeleter.DeleteItemsByUserID(ctx, userID); err != nil {
		return models.OrderID(0), fmt.Errorf("delete items: %w", err)
	}

	return orderID, nil
}
