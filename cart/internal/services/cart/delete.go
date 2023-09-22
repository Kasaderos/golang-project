package cart

import (
	"context"
	"route256/cart/internal/models"
)

type ItemDeleter interface {
	DeleteItem(ctx context.Context, userID models.UserID, SKU models.SKU) error
	DeleteItemsByUserID(ctx context.Context, userID models.UserID) error
}

type DeleteService struct {
	itemDeleter ItemDeleter
}

func NewDeleteService(itemDeleter ItemDeleter) *DeleteService {
	return &DeleteService{
		itemDeleter: itemDeleter,
	}
}

func (c DeleteService) DeleteItem(ctx context.Context, userID models.UserID, sku models.SKU) error {
	return c.itemDeleter.DeleteItem(ctx, userID, sku)
}

func (c DeleteService) Clear(ctx context.Context, userID models.UserID) error {
	return c.itemDeleter.DeleteItemsByUserID(ctx, userID)
}
