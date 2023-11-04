package order

import (
	"context"
	"errors"
	"route256/loms/internal/models"
	"route256/loms/internal/services/order/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCancelService_CancelOrder(t *testing.T) {
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
	errReserveCanceller := errors.New("reserve cancel failed")
	orderProviderErrMock := mock.NewOrderProviderMock(t).
		GetOrderByIDMock.
		Expect(ctx, order.ID).
		Return(nil, errOrderProvider)
	orderProviderSuccessMock := mock.NewOrderProviderMock(t).
		GetOrderByIDMock.
		Expect(ctx, order.ID).
		Return(order, nil)

	reserveCancellerErrMock := mock.NewReserveCancellerMock(t).
		ReserveCancelMock.
		Expect(ctx, order.Items).
		Return(errReserveCanceller)

	reserveCancellerSuccessMock := mock.NewReserveCancellerMock(t).
		ReserveCancelMock.
		Expect(ctx, order.Items).
		Return(nil)

	orderStatusSetterCancelMock := mock.NewOrderStatusSetterMock(t).
		SetStatusMock.
		Expect(ctx, order.ID, models.StatusCancelled).
		Return(nil)

	statusNotifierSuccessMock := mock.NewStatusNotifierMock(t).
		NotifyOrderStatusMock.
		Expect(ctx, order.ID, models.StatusCancelled).
		Return(nil)

	tests := []struct {
		name string

		orderProvide      OrderProvider
		reserveCanceller  ReserveCanceller
		orderStatusSetter OrderStatusSetter
		statusNotifier    StatusNotifier

		orderID models.OrderID

		errAssert require.ErrorAssertionFunc
	}{
		{
			name:         "get order failed",
			orderProvide: orderProviderErrMock,
			orderID:      order.ID,
			errAssert:    require.Error,
		},
		{
			name:             "reserve cancel err",
			orderProvide:     orderProviderSuccessMock,
			reserveCanceller: reserveCancellerErrMock,
			orderID:          order.ID,
			errAssert:        require.Error,
		},
		{
			name:              "cancelled success",
			orderProvide:      orderProviderSuccessMock,
			reserveCanceller:  reserveCancellerSuccessMock,
			orderStatusSetter: orderStatusSetterCancelMock,
			statusNotifier:    statusNotifierSuccessMock,

			orderID:   order.ID,
			errAssert: require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			usc := &CancelService{
				orderProvide:      tt.orderProvide,
				reserveCanceller:  tt.reserveCanceller,
				orderStatusSetter: tt.orderStatusSetter,
				statusNotifier:    tt.statusNotifier,
			}

			err := usc.CancelOrder(ctx, tt.orderID)
			tt.errAssert(t, err)
		})
	}

	t.Cleanup(func() {
		orderProviderErrMock.MinimockFinish()
		reserveCancellerErrMock.MinimockFinish()
		orderProviderSuccessMock.MinimockFinish()
		reserveCancellerSuccessMock.MinimockFinish()
		orderStatusSetterCancelMock.MinimockFinish()
		statusNotifierSuccessMock.MinimockFinish()
	})
}
