package cart

import (
	"context"
	"route256/cart/internal/models"
)

type ItemsDeleter interface {
	DeleteItemsByUserID(ctx context.Context, userID models.UserID) error
}

type ClearService struct {
	itemsDeleter ItemsDeleter
}

func NewClearService(itemsDeleter ItemsDeleter) *ClearService {
	return &ClearService{
		itemsDeleter: itemsDeleter,
	}
}

func (c ClearService) Clear(ctx context.Context, userID models.UserID) error {
	return c.itemsDeleter.DeleteItemsByUserID(ctx, userID)
}
