package repository

import (
	"context"
	"fmt"
	"route256/loms/internal/models"
	oms "route256/loms/internal/usecase/OMS"
	"sync"
	"time"
)

const OrderExpiration = time.Minute * 10

type omsRepository struct {
	mu sync.RWMutex
	db map[OrderID]*Order
}

var _ oms.OMSRepository = (*omsRepository)(nil)

func NewOMSRepostiory() *omsRepository {
	return &omsRepository{}
}

func (r *omsRepository) CreateOrder(ctx context.Context, order models.Order) error {
	rOrder := &Order{
		ID:     OrderID(order.ID),
		UserID: UserID(order.UserID),
		Status: order.Status.String(),
		Items:  make([]ItemOrderInfo, 0, len(order.Items)),
	}
	for _, item := range order.Items {
		rOrder.Items = append(rOrder.Items, ItemOrderInfo{
			SKU:   SKU(item.SKU),
			Count: item.Count,
		})
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.db[rOrder.ID]
	if ok {
		return fmt.Errorf("repo: order exist %v", order.ID)
	}
	r.db[rOrder.ID] = rOrder

	return nil
}

func (r *omsRepository) GetOrderByID(ctx context.Context, orderID models.OrderID) (*models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.db[OrderID(orderID)]
	if !ok {
		return nil, fmt.Errorf("repo: order not found %v", orderID)
	}

	var result models.Order
	order.Scan(&result)

	return &result, nil
}

func (r *omsRepository) SetStatus(ctx context.Context, orderID models.OrderID, status models.Status) error {
	r.mu.Lock()
	order, ok := r.db[OrderID(orderID)]
	if !ok {
		return fmt.Errorf("repo: order not found %v", orderID)
	}
	order.Status = status.String()
	r.mu.Unlock()

	return nil
}

func (r *omsRepository) ListExpiredOrders(ctx context.Context, limit uint32) ([]models.OrderID, error) {
	expiredOrderIDs := make([]models.OrderID, 0)
	now := time.Now()

	r.mu.RLock()
	for id, order := range r.db {
		if len(expiredOrderIDs) > int(limit) {
			return expiredOrderIDs, nil
		}
		if now.Sub(order.CreatedAt) > OrderExpiration {
			expiredOrderIDs = append(expiredOrderIDs, models.OrderID(id))
		}
	}
	r.mu.RUnlock()

	return expiredOrderIDs, nil
}
