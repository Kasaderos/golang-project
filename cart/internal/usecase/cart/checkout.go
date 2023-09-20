package cart

import (
	"context"
	"route256/cart/internal/models"
)

func (usc cartUsecase) Checkout(ctx context.Context, userID models.UserID) (models.OrderID, error) {
	items, err := usc.CartRepository.GetItemsByUserID(ctx, userID)
	if err != nil {
		return models.OrderID(0), err
	}

	orderID, err := usc.LOMSService.CreateOrder(ctx, userID, items)
	if err != nil {
		return models.OrderID(0), err
	}

	if err := usc.CartRepository.DeleteItemsByUserID(ctx, userID); err != nil {
		return models.OrderID(0), err
	}

	return orderID, nil
}
