package repository

import (
	"context"
	"route256/loms/internal/models"
	oms "route256/loms/internal/usecase/OMS"
	"time"
)

const OrderExpiration = time.Minute * 10

type omsRepository struct {
}

var _ oms.OMSRepository = (*omsRepository)(nil)

func NewOMSRepostiory() *omsRepository {
	return &omsRepository{}
}

func (r *omsRepository) CreateOrder(ctx context.Context, order models.Order) error {
	return nil
}

func (r *omsRepository) GetOrderByID(ctx context.Context, orderID models.OrderID) (*models.Order, error) {
	return nil, nil
}

func (r *omsRepository) SetStatus(ctx context.Context, orderID models.OrderID, status models.Status) error {
	return nil
}

func (r *omsRepository) ListExpiredOrders(ctx context.Context, limit uint32) ([]models.OrderID, error) {
	return nil, nil
}
