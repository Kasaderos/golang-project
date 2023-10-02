package repository

import (
	"context"
	"math/rand"
	"route256/loms/internal/models"
	"time"
)

const OrderExpiration = time.Minute * 10

type OrdersRepository struct {
}

func NewOMSRepostiory() *OrdersRepository {
	return &OrdersRepository{}
}

func (r *OrdersRepository) CreateOrder(ctx context.Context, order models.Order) (models.OrderID, error) {
	return models.OrderID(rand.Int() % 1000), nil
}

func (r *OrdersRepository) GetOrderByID(ctx context.Context, orderID models.OrderID) (*models.Order, error) {
	return &models.Order{
		ID: orderID,
	}, nil
}

func (r *OrdersRepository) SetStatus(ctx context.Context, orderID models.OrderID, status models.Status) error {
	return nil
}

func (r *OrdersRepository) ListExpiredOrders(ctx context.Context, limit uint32) ([]models.OrderID, error) {
	return []models.OrderID{}, nil
}
