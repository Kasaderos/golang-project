package oms

// import (
// 	"context"
// 	"route256/loms/internal/pkg/models"
// )

// type OrdersProvider interface {
// 	GetOrderByID(ctx context.Context, orderID int64) (models.Order, error)
// }

// type StocksReserveRemover interface {
// 	ReserveRemover(ctx context.Context, orderID int64) error
// }

// type OrderService struct {
// 	orderStorage   OrdersStorage
// 	stocksReserver StocksReserver
// }

// func NewOrderService(orderStorage OrdersStorage, reserver StocksReserver) *OrderService {
// 	return &OrderService{
// 		orderStorage:   orderStorage,
// 		stocksReserver: reserver,
// 	}
// }

// func (o *OrderService) Create(ctx context.Context, order models.Order) (orderID string, err error) {
// }
