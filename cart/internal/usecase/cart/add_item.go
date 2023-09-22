package cart

import (
	"context"
	"fmt"
	"route256/cart/internal/models"
)

func (usc cartUsecase) AddItem(ctx context.Context, userID models.UserID, sku models.SKU, count uint16) error {
	_, _, err := usc.ProductService.GetProductInfo(ctx, sku)
	if err != nil {
		return fmt.Errorf("product service: %w", err)
	}

	stockCount, err := usc.LOMSService.GetStock(ctx, sku)
	if err != nil {
		return err
	}

	if uint64(count) > stockCount {
		return fmt.Errorf("add item: not enough stocks %d > %d", count, stockCount)
	}

	return usc.CartRepository.AddItem(ctx, userID, models.CartItem{
		SKU:   sku,
		Count: count,
	})
}
