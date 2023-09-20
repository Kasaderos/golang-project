package cart

import (
	"context"
	"route256/cart/internal/models"
)

func (usc cartUsecase) DeleteItem(ctx context.Context, userID models.UserID, sku models.SKU) error {
	return usc.CartRepository.DeleteItem(ctx, userID, sku)
}

func (usc cartUsecase) Clear(ctx context.Context, userID models.UserID) error {
	return usc.CartRepository.DeleteItemsByUserID(ctx, userID)
}
