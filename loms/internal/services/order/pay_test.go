package order

import (
	"context"
	"errors"
	"route256/loms/internal/models"
	"route256/loms/internal/services/order/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPayService_MarkAsPaid(t *testing.T) {
	ctx := context.Background()

	var order = &models.Order{
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

	errOrderProvider := errors.New("get order failed")
	orderProviderErrMock := mock.NewOrderProviderMock(t).
		GetOrderByIDMock.
		Expect(ctx, order.ID).
		Return(nil, errOrderProvider)
	orderProviderSuccessMock := mock.NewOrderProviderMock(t).
		GetOrderByIDMock.
		Expect(ctx, order.ID).
		Return(order, nil)

	errReserveRemove := errors.New("reserve remove failed")
	reserveRemoverErrMock := mock.NewReserveRemoverMock(t).
		ReserveRemoveMock.
		Expect(ctx, order.Items).
		Return(errReserveRemove)
	reserveRemoverSuccessMock := mock.NewReserveRemoverMock(t).
		ReserveRemoveMock.
		Expect(ctx, order.Items).
		Return(nil)

	orderStatusSetterPaidMock := mock.NewOrderStatusSetterMock(t).
		SetStatusMock.
		Expect(ctx, order.ID, models.StatusPaid).
		Return(nil)

	tests := []struct {
		name string

		orderProvider     OrderProvider
		reserveRemover    ReserveRemover
		orderStatusSetter OrderStatusSetter

		orderID models.OrderID

		errAssert require.ErrorAssertionFunc
	}{
		{
			name:          "order provide down",
			orderProvider: orderProviderErrMock,
			orderID:       order.ID,

			errAssert: require.Error,
		},
		{
			name:           "order provide down",
			orderProvider:  orderProviderSuccessMock,
			reserveRemover: reserveRemoverErrMock,
			orderID:        order.ID,

			errAssert: require.Error,
		},
		{
			name:              "paid successfully",
			orderProvider:     orderProviderSuccessMock,
			reserveRemover:    reserveRemoverSuccessMock,
			orderStatusSetter: orderStatusSetterPaidMock,
			orderID:           order.ID,

			errAssert: require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			usc := &PayService{
				orderProvider:     tt.orderProvider,
				reserveRemover:    tt.reserveRemover,
				orderStatusSetter: tt.orderStatusSetter,
			}
			err := usc.MarkAsPaid(ctx, tt.orderID)

			tt.errAssert(t, err)
		})
	}

	t.Cleanup(func() {
		orderProviderErrMock.MinimockFinish()
		orderProviderSuccessMock.MinimockFinish()
		reserveRemoverErrMock.MinimockFinish()
	})
}
