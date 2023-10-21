package cart

import (
	"context"
	"errors"
	"route256/cart/internal/models"
	"route256/cart/internal/services"
	"route256/cart/internal/services/cart/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCart_ItemAdd(t *testing.T) {
	ctx := context.Background()
	const (
		sku             = models.SKU(0)
		userID          = models.UserID(1)
		productName     = "airpods"
		productPrice    = 100
		productCount    = 2
		bigProductCount = 10
		stockCount      = 3
	)

	var cartItemSuccessCase = models.CartItem{
		SKU:   sku,
		Name:  productName,
		Count: productCount,
		Price: productPrice,
	}

	errNotFound := errors.New("not found")

	productProviderErrMock := mock.NewProductProviderMock(t).
		GetProductInfoMock.
		Expect(ctx, sku).
		Return("", 0, errNotFound)

	productProviderSuccessMock := mock.NewProductProviderMock(t).
		GetProductInfoMock.
		Expect(ctx, sku).
		Return(productName, productPrice, nil)

	stockProviderErrMock := mock.NewStockProviderMock(t).
		GetStockMock.
		Expect(ctx, sku).
		Return(0, errNotFound)

	stockProviderSuccessMock := mock.NewStockProviderMock(t).
		GetStockMock.
		Expect(ctx, sku).
		Return(stockCount, nil)

	itemAdderSuccessMock := mock.NewItemAdderMock(t).
		AddItemMock.
		Expect(ctx, userID, cartItemSuccessCase).
		Return(nil)

	tests := []struct {
		name                string
		productProviderMock ProductProvider
		stockProviderMock   StockProvider
		itemAdderMock       ItemAdder

		userID models.UserID
		sku    models.SKU
		count  uint16

		wantServiceError error
		errAssert        require.ErrorAssertionFunc
	}{
		{
			name:                "not found product by sku",
			productProviderMock: productProviderErrMock,
			stockProviderMock:   stockProviderSuccessMock,
			itemAdderMock:       itemAdderSuccessMock,

			userID: userID,
			sku:    sku,
			count:  productCount,

			errAssert: require.Error,
		},
		{
			name:                "not found stock by sku",
			productProviderMock: productProviderSuccessMock,
			stockProviderMock:   stockProviderErrMock,
			itemAdderMock:       itemAdderSuccessMock,

			userID: userID,
			sku:    sku,
			count:  productCount,

			errAssert: require.Error,
		},
		{
			name:                "product count > stock count",
			productProviderMock: productProviderSuccessMock,
			stockProviderMock:   stockProviderSuccessMock,
			itemAdderMock:       itemAdderSuccessMock,

			userID: userID,
			sku:    sku,
			count:  bigProductCount,

			wantServiceError: services.ErrNotEnoughStocks,
			errAssert:        require.Error,
		},
		{
			name:                "success",
			productProviderMock: productProviderSuccessMock,
			stockProviderMock:   stockProviderSuccessMock,
			itemAdderMock:       itemAdderSuccessMock,

			userID: userID,
			sku:    sku,
			count:  productCount,

			errAssert: require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := NewAddService(AddDeps{
				ProductProvider: tt.productProviderMock,
				StockProvider:   tt.stockProviderMock,
				ItemAdder:       tt.itemAdderMock,
			})
			err := s.AddItem(
				ctx,
				tt.userID,
				tt.sku,
				tt.count,
			)
			tt.errAssert(t, err)

			if tt.wantServiceError != nil {
				var serviceErr *services.CartServiceError
				require.ErrorAs(t, err, &serviceErr)
			}
		})
	}
	t.Cleanup(func() {
		productProviderErrMock.MinimockFinish()
		productProviderSuccessMock.MinimockFinish()

		stockProviderErrMock.MinimockFinish()
		stockProviderSuccessMock.MinimockFinish()

		itemAdderSuccessMock.MinimockFinish()
	})
}
