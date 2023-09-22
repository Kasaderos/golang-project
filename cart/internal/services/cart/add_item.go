package cart

import (
	"context"
	"fmt"
	"route256/cart/internal/models"
)

type ProductProvider interface {
	GetProductInfo(cxt context.Context, sku models.SKU) (name string, price uint32, err error)
}

type StockProvider interface {
	GetStock(ctx context.Context, sku models.SKU) (count uint64, err error)
}

type ItemAdder interface {
	AddItem(ctx context.Context, userID models.UserID, item models.CartItem) error
}

type AddService struct {
	productProvider ProductProvider
	stockProvider   StockProvider
	itemAdder       ItemAdder
}

type AddDeps struct {
	ProductProvider
	StockProvider
	ItemAdder
}

func NewAddService(d AddDeps) *AddService {
	return &AddService{
		productProvider: d.ProductProvider,
		stockProvider:   d.StockProvider,
		itemAdder:       d.ItemAdder,
	}
}

func (c AddService) AddItem(ctx context.Context, userID models.UserID, sku models.SKU, count uint16) error {
	_, _, err := c.productProvider.GetProductInfo(ctx, sku)
	if err != nil {
		return fmt.Errorf("product service: %w", err)
	}

	stockCount, err := c.stockProvider.GetStock(ctx, sku)
	if err != nil {
		return err
	}

	if uint64(count) > stockCount {
		return fmt.Errorf("add item: not enough stocks %d > %d", count, stockCount)
	}

	return c.itemAdder.AddItem(ctx, userID, models.CartItem{
		SKU:   sku,
		Count: count,
	})
}
