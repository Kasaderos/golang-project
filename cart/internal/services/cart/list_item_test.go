package cart

import (
	"context"
	"errors"
	"route256/cart/internal/models"
	"route256/cart/internal/services/cart/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCart_ListItem(t *testing.T) {
	ctx := context.Background()
	const (
		userID     = models.UserID(1)
		totalPrice = 60
	)

	var cartItems = []models.CartItem{
		{
			SKU:   0,
			Name:  "product1",
			Count: 1,
			Price: 10,
		},
		{
			SKU:   1,
			Name:  "product2",
			Count: 2,
			Price: 10,
		},
		{
			SKU:   2,
			Name:  "product3",
			Count: 3,
			Price: 10,
		},
	}

	errNotFound := errors.New("not found")
	errClientError := errors.New("client error")

	itemsProviderErrMock := mock.NewItemsProviderMock(t).
		GetItemsByUserIDMock.
		Expect(ctx, userID).
		Return(nil, errNotFound)

	itemsProviderSuccessMock := mock.NewItemsProviderMock(t).
		GetItemsByUserIDMock.
		Expect(ctx, userID).
		Return(cartItems, nil)

	itemsProviderNoItemsMock := mock.NewItemsProviderMock(t).
		GetItemsByUserIDMock.
		Expect(ctx, userID).
		Return(nil, nil)

	productProviderErrMock := mock.NewProductProviderMock(t).
		GetProductInfoMock.
		// not expecting ctx, because ctx will be changed in workerpool
		// Expect(ctx, sku).
		Return("", 0, errClientError)

	productProviderSuccessMock := mock.NewProductProviderMock(t)
	for _, item := range cartItems {
		productProviderSuccessMock.GetProductInfoMock.
			// not expecting ctx, because ctx will be changed in workerpool
			// Expect(ctx, sku).
			Return(item.Name, item.Price, nil)
	}

	tests := []struct {
		name                string
		productProviderMock ProductProvider
		itemsProviderMock   ItemsProvider

		userID     models.UserID
		maxWorkers int

		wantTotalPrice uint32
		wantItems      []models.CartItem
		errAssert      require.ErrorAssertionFunc
	}{
		{
			name:              "no items",
			itemsProviderMock: itemsProviderNoItemsMock,

			userID:     userID,
			maxWorkers: 1,

			errAssert: require.NoError,
		},
		{
			name:              "failed get items",
			itemsProviderMock: itemsProviderErrMock,

			userID:     userID,
			maxWorkers: 1,

			errAssert: require.Error,
		},
		{
			name:                "product provider down",
			itemsProviderMock:   itemsProviderSuccessMock,
			productProviderMock: productProviderErrMock,

			userID:     userID,
			maxWorkers: 2,

			errAssert: require.Error,
		},
		{
			name:                "success",
			itemsProviderMock:   itemsProviderSuccessMock,
			productProviderMock: productProviderSuccessMock,

			userID:         userID,
			maxWorkers:     2,
			wantTotalPrice: totalPrice,
			wantItems:      cartItems,

			errAssert: require.NoError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := NewListItemService(
				tt.itemsProviderMock,
				tt.productProviderMock,
				tt.maxWorkers,
			)
			totalPrice, items, err := s.ListItem(
				ctx,
				tt.userID,
			)
			tt.errAssert(t, err)
			require.Equal(t, tt.wantTotalPrice, totalPrice, "actual", totalPrice)
			require.Equal(t, tt.wantItems, items)
		})
	}

	t.Cleanup(func() {
		itemsProviderErrMock.MinimockFinish()
		itemsProviderNoItemsMock.MinimockFinish()
		itemsProviderSuccessMock.MinimockFinish()

		productProviderErrMock.MinimockFinish()
		productProviderSuccessMock.MinimockFinish()
	})
}
