package cart

import (
	"context"
	"route256/cart/internal/models"
)

type ItemsProvider interface {
	GetItemsByUserID(ctx context.Context, userID models.UserID) ([]models.CartItem, error)
}

type ListService struct {
	itemsProvider   ItemsProvider
	productProvider ProductProvider
}

func NewListService(itemPr ItemsProvider, productPr ProductProvider) *ListService {
	return &ListService{
		itemsProvider:   itemPr,
		productProvider: productPr,
	}
}

func (c ListService) ListItem(
	ctx context.Context,
	userID models.UserID,
) (totalPrice uint32, items []models.CartItem, err error) {
	items, err = c.itemsProvider.GetItemsByUserID(ctx, userID)
	if err != nil {
		return 0, nil, err
	}

	for i, item := range items {
		name, price, err := c.productProvider.GetProductInfo(ctx, item.SKU)
		if err != nil {
			return 0, nil, err
		}

		items[i].Name = name
		items[i].Price = price

		totalPrice += price
	}

	return totalPrice, items, nil
}
