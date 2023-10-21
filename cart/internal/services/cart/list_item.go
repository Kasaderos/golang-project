package cart

import (
	"context"
	"route256/cart/internal/models"
	"route256/cart/pkg/workerpool"
	"sync/atomic"
)

type ItemsProvider interface {
	GetItemsByUserID(ctx context.Context, userID models.UserID) ([]models.CartItem, error)
}

type ListService struct {
	itemsProvider   ItemsProvider
	productProvider ProductProvider
	maxWorkers      int
}

func NewListItemService(
	itemPr ItemsProvider,
	productPr ProductProvider,
	maxWorkers int,
) *ListService {
	return &ListService{
		itemsProvider:   itemPr,
		productProvider: productPr,
		maxWorkers:      maxWorkers,
	}
}

func (c ListService) ListItem(
	ctx context.Context,
	userID models.UserID,
) (uint32, []models.CartItem, error) {
	// assume this user don't have many items
	items, err := c.itemsProvider.GetItemsByUserID(ctx, userID)
	if err != nil {
		return 0, nil, err
	}

	// assume that the prices and names often change
	// another way of solution
	// we might cache price and name by skus
	// and just update cache in separate goroutine
	totalPrice := uint32(0)
	wp, ctx := workerpool.New(ctx, c.maxWorkers)
	for i, item := range items {
		i := i
		item := item
		wp.Run(func() error {
			name, price, err := c.productProvider.GetProductInfo(ctx, item.SKU)
			if err != nil {
				return err
			}

			items[i].Name = name
			items[i].Price = price

			atomic.AddUint32(&totalPrice, price*uint32(item.Count))

			return nil
		})
	}
	if err := wp.Wait(); err != nil {
		return 0, nil, err
	}

	return totalPrice, items, nil
}
