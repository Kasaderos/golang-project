package order

import (
	"context"
	"errors"
	"route256/loms/internal/models"
	"route256/loms/internal/services/order/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateService_CreateOrder(t *testing.T) {
	ctx := context.Background()

	var order = models.Order{
		ID:     0,
		UserID: 100500,
		Status: models.StatusNew,
		Items: []models.ItemOrderInfo{
			{
				SKU:   1,
				Count: 10,
			},
			{
				SKU:   2,
				Count: 20,
			},
		},
	}

	errOrderCreate := errors.New("create order failed")
	errStockReserve := errors.New("stock reserve failed")

	orderCreatorErrMock := mock.NewOrderCreatorMock(t).
		CreateOrderMock.
		Expect(ctx, order).
		Return(models.OrderID(-1), errOrderCreate)
	orderCreatorSuccessMock := mock.NewOrderCreatorMock(t).
		CreateOrderMock.
		Expect(ctx, order).
		Return(order.ID, nil)

	stockReserverErrMock := mock.NewStocksReserverMock(t).
		ReserveStocksMock.
		Expect(ctx, order.Items).
		Return(errStockReserve)

	stockReserverSuccessMock := mock.NewStocksReserverMock(t).
		ReserveStocksMock.
		Expect(ctx, order.Items).
		Return(nil)

	orderStatusSetterFailMock := mock.NewOrderStatusSetterMock(t).
		SetStatusMock.
		Expect(ctx, order.ID, models.StatusFailed).
		Return(nil)

	orderStatusSetterSuccessMock := mock.NewOrderStatusSetterMock(t).
		SetStatusMock.
		Expect(ctx, order.ID, models.StatusAwaitingPayment).
		Return(nil)

	tests := []struct {
		name string

		orderCreator      OrderCreator
		stocksReserver    StocksReserver
		orderStatusSetter OrderStatusSetter

		userID      models.UserID
		items       []models.ItemOrderInfo
		wantOrderID models.OrderID

		errAssert require.ErrorAssertionFunc
	}{
		{
			name:         "create order failed",
			orderCreator: orderCreatorErrMock,
			userID:       order.UserID,
			items:        order.Items,
			wantOrderID:  models.OrderID(-1),
			errAssert:    require.Error,
		},
		{
			name:              "reserve stock failed",
			orderCreator:      orderCreatorSuccessMock,
			stocksReserver:    stockReserverErrMock,
			orderStatusSetter: orderStatusSetterFailMock,
			userID:            order.UserID,
			items:             order.Items,
			wantOrderID:       models.OrderID(-1),
			errAssert:         require.Error,
		},
		{
			name:              "order created successfully",
			orderCreator:      orderCreatorSuccessMock,
			stocksReserver:    stockReserverSuccessMock,
			orderStatusSetter: orderStatusSetterSuccessMock,
			userID:            order.UserID,
			items:             order.Items,
			wantOrderID:       order.ID,
			errAssert:         require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usc := &CreateService{
				orderCreator:      tt.orderCreator,
				stocksReserver:    tt.stocksReserver,
				orderStatusSetter: tt.orderStatusSetter,
			}
			got, err := usc.CreateOrder(ctx, tt.userID, tt.items)

			tt.errAssert(t, err)
			require.Equal(t, tt.wantOrderID, got)
		})
	}
	t.Cleanup(func() {
		orderCreatorErrMock.MinimockFinish()
		orderCreatorSuccessMock.MinimockFinish()
		stockReserverErrMock.MinimockFinish()
		stockReserverSuccessMock.MinimockFinish()
		orderStatusSetterFailMock.MinimockFinish()
	})
}
