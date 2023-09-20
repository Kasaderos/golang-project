package oms

import (
	"context"
	"math/rand"
	"route256/loms/internal/models"
	"route256/loms/internal/usecase"
)

func (usc *omsUsecase) CreateOrder(
	ctx context.Context,
	userID models.UserID,
	info usecase.CreateOrderInfo,
) (models.OrderID, error) {
	var (
		OrderID = models.OrderID(rand.Int() % 1000)
		order   = models.Order{
			ID:     OrderID,
			UserID: userID,
			Items:  info.Items,
		}
	)

	if err := usc.OMSRepository.CreateOrder(ctx, order); err != nil {
		return models.OrderID(-1), err
	}

	if err := usc.WMSRepository.ReserveStocks(ctx, userID, info.Items); err != nil {
		return models.OrderID(-1), err
	}

	return OrderID, nil
}
