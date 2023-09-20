package cart

import (
	"context"
	"route256/cart/internal/models"
)

func (usc cartUsecase) ListItem(
	ctx context.Context,
	userID models.UserID,
) (totalPrice uint32, items []models.CartItem, err error) {
	items, err = usc.CartRepository.GetItemsByUserID(ctx, userID)
	if err != nil {
		return 0, nil, err
	}

	for i, item := range items {
		name, price, err := usc.ProductService.GetProductInfo(ctx, item.SKU)
		if err != nil {
			return 0, nil, err
		}

		items[i].Name = name
		items[i].Price = price

		totalPrice += price
	}

	return totalPrice, items, nil
}
