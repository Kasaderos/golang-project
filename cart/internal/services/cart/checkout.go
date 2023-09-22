package cart

import (
	"context"
	"route256/cart/internal/models"
)

type OrderCreator interface {
	CreateOrder(ctx context.Context, userID models.UserID, items []models.CartItem) (models.OrderID, error)
}

type CheckoutService struct {
	orderCreator  OrderCreator
	itemsProvider ItemsProvider
	itemsDeleter  ItemDeleter
}

type CheckoutDeps struct {
	OrderCreator
	ItemsProvider
	ItemDeleter
}

func NewCheckoutService(d CheckoutDeps) *CheckoutService {
	return &CheckoutService{
		orderCreator:  d.OrderCreator,
		itemsProvider: d.ItemsProvider,
		itemsDeleter:  d.ItemDeleter,
	}
}

func (c CheckoutService) Checkout(ctx context.Context, userID models.UserID) (models.OrderID, error) {
	items, err := c.itemsProvider.GetItemsByUserID(ctx, userID)
	if err != nil {
		return models.OrderID(0), err
	}

	orderID, err := c.orderCreator.CreateOrder(ctx, userID, items)
	if err != nil {
		return models.OrderID(0), err
	}

	if err := c.itemsDeleter.DeleteItemsByUserID(ctx, userID); err != nil {
		return models.OrderID(0), err
	}

	return orderID, nil
}
