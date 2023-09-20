package cart

import (
	"context"
	"route256/cart/internal/models"
	"route256/cart/internal/usecase"
)

type (
	CartRepository interface {
		CartItemAdder
		CartItemDeleter
		CartItemsProvider
	}

	CartItemAdder interface {
		AddItem(ctx context.Context, userID models.UserID, item models.CartItem) error
	}

	CartItemDeleter interface {
		DeleteItem(ctx context.Context, userID models.UserID, SKU models.SKU) error
		DeleteItemsByUserID(ctx context.Context, userID models.UserID) error
	}

	CartItemsProvider interface {
		GetItemsByUserID(ctx context.Context, userID models.UserID) ([]models.CartItem, error)
	}

	LOMSService interface {
		StockProvider
		OrderCreator
	}

	OrderCreator interface {
		CreateOrder(ctx context.Context, userID models.UserID, items []models.CartItem) (models.OrderID, error)
	}

	StockProvider interface {
		GetStock(ctx context.Context, sku models.SKU) (count uint64, err error)
	}

	ProductService interface {
		GetProductInfo(cxt context.Context, sku models.SKU) (name string, price uint32, err error)
	}
)

type Deps struct {
	CartRepository
	LOMSService
	ProductService
}

type cartUsecase struct {
	Deps
}

var _ usecase.CartService = (*cartUsecase)(nil)

func NewCartUsecase(d Deps) *cartUsecase {
	return &cartUsecase{
		Deps: d,
	}
}
