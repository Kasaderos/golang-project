package cart

import (
	"context"
	"errors"
	"route256/cart/internal/models"
)

func (usc cartUsecase) AddItem(ctx context.Context, userID models.UserID, sku models.SKU, count uint16) error {
	_, _, err := usc.ProductService.GetProductInfo(ctx, sku)
	if err != nil {
		return err
	}

	stockCount, err := usc.LOMSService.GetStock(ctx, sku)
	if err != nil {
		return err
	}

	if uint64(count) > stockCount {
		return errors.New("todo")
	}

	return usc.CartRepository.AddItem(ctx, userID, models.CartItem{
		SKU:   sku,
		Count: count,
	})
}
