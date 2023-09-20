package oms

import (
	"context"
	"route256/loms/internal/models"
)

func (usc *omsUsecase) GetOrderInfo(
	ctx context.Context,
	orderID models.OrderID,
) (*models.Order, error) {
	return usc.OMSRepository.GetOrderByID(ctx, orderID)
}

func (usc *omsUsecase) MarkOrderAsPaid(
	ctx context.Context,
	orderID models.OrderID,
) error {
	order, err := usc.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	if err := usc.WMSRepository.ReserveRemove(ctx, order.UserID, order.Items); err != nil {
		return err
	}

	return usc.OMSRepository.SetStatus(ctx, orderID, models.StatusPaid)
}
