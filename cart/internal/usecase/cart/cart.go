package cart

import (
	"context"
	"errors"
	"route256/cart/internal/models"
)

type (
	CartRepository interface {
		AddItem(ctx context.Context, userID models.UserID, item models.CartItem) error
		GetItemsByUserID(ctx context.Context, userID models.UserID) ([]models.CartItem, error)
		DeleteItem(ctx context.Context, userID models.UserID, SKU models.SKU) error
		DeleteItemsByUserID(ctx context.Context, userID models.UserID) error
	}

	ProductService interface {
		GetProductInfo(cxt context.Context, sku models.SKU) (name string, price uint32, err error)
	}

	LOMSService interface {
		CreateOrder(ctx context.Context, userID models.UserID, items []models.CartItem) error
		GetStock(ctx context.Context, sku models.SKU) (count uint64, err error)
	}
)

type Deps struct {
	CartRepository
	ProductService
	LOMSService
}

type cartUsecase struct {
	Deps
}

func NewCartUsecase(d Deps) *cartUsecase {
	return &cartUsecase{
		Deps: d,
	}
}

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

func (usc cartUsecase) ListItem(
	ctx context.Context,
	userID models.UserID,
	sku models.SKU,
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

func (usc cartUsecase) Checkout(ctx context.Context, userID models.UserID) error {
	items, err := usc.CartRepository.GetItemsByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if err := usc.LOMSService.CreateOrder(ctx, userID, items); err != nil {
		return err
	}

	return usc.CartRepository.DeleteItemsByUserID(ctx, userID)
}

func (usc cartUsecase) DeleteItem(ctx context.Context, userID models.UserID, sku models.SKU) error {
	return usc.CartRepository.DeleteItem(ctx, userID, sku)
}

func (usc cartUsecase) Clear(ctx context.Context, userID models.UserID) error {
	return usc.CartRepository.DeleteItemsByUserID(ctx, userID)
}
