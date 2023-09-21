package oms

import (
	"context"
	"route256/loms/internal/models"
)

func (usc *omsUsecase) CancelOrder(
	ctx context.Context,
	orderID models.OrderID,
) error {
	order, err := usc.OMSRepository.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	if err := usc.WMSRepository.ReserveCancel(ctx, order.UserID, order.Items); err != nil {
		return err
	}

	return usc.OMSRepository.SetStatus(ctx, orderID, models.StatusCancelled)
}
