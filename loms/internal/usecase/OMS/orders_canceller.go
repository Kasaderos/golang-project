package oms

import (
	"context"
	"log"
	"time"
)

const (
	ordersCheckDuration   = time.Minute
	listActiveOrdersLimit = 100
)

func (usc omsUsecase) CancelNotPaidOrdersBackground(ctx context.Context) {
	ticker := time.NewTicker(ordersCheckDuration)
	for {
		select {
		case <-ticker.C:
			usc.cancelNotPaidOrders(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (usc omsUsecase) cancelNotPaidOrders(ctx context.Context) {
	expiredOrderIDs, err := usc.OMSRepository.ListExpiredOrders(ctx, listActiveOrdersLimit)
	if err != nil {
		log.Printf("can't list orders: %v", err)
		return
	}

	for _, id := range expiredOrderIDs {
		if err = usc.CancelOrder(ctx, id); err != nil {
			log.Printf("while cancelling expired order: order_id %v: %v", id, err)
		}
	}
}
