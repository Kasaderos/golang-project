package repository

import (
	"context"
	"route256/loms/internal/models"
	"time"
)

const OrderExpiration = time.Minute * 10

type omsRepository struct {
}

func NewOMSRepostiory() *omsRepository {
	return &omsRepository{}
}

func (r *omsRepository) CreateOrder(ctx context.Context, order models.Order) error {
	return nil
}

func (r *omsRepository) GetOrderByID(ctx context.Context, orderID models.OrderID) (*models.Order, error) {
	return &models.Order{
		ID: orderID,
	}, nil
}

func (r *omsRepository) SetStatus(ctx context.Context, orderID models.OrderID, status models.Status) error {
	return nil
}

func (r *omsRepository) ListExpiredOrders(ctx context.Context, limit uint32) ([]models.OrderID, error) {
	return []models.OrderID{}, nil
}
