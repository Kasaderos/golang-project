package postgres

import (
	"context"
	"fmt"
	"log"

	"route256/loms/internal/models"

	sqlc "route256/loms/internal/repository/postgres/order"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrdersRepository struct {
	dbpool *pgxpool.Pool
}

func NewOrdersRepostiory(dbpool *pgxpool.Pool) *OrdersRepository {
	return &OrdersRepository{
		dbpool: dbpool,
	}
}

func (r *OrdersRepository) CreateOrder(
	ctx context.Context,
	order models.Order,
) (models.OrderID, error) {
	tx, err := r.dbpool.Begin(ctx)
	if err != nil {
		return models.OrderID(0), err
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Println(err)
		}
	}()

	q := sqlc.New(r.dbpool)
	q = q.WithTx(tx)

	id, err := q.CreateOrder(ctx, sqlc.CreateOrderParams{
		UserID: pgtype.Int8{
			Int64: int64(order.UserID),
			Valid: true,
		},
		Status: pgtype.Text{
			String: order.Status.String(),
			Valid:  true,
		},
	})
	if err != nil {
		return models.OrderID(0), fmt.Errorf("create order: %w", err)
	}

	for _, item := range order.Items {
		if err := q.AddOrderItem(ctx, sqlc.AddOrderItemParams{
			OrderID: pgtype.Int8{
				Int64: id,
				Valid: true,
			},
			Sku: pgtype.Int8{
				Int64: int64(item.SKU),
				Valid: true,
			},
			Count: pgtype.Int4{
				Int32: int32(item.Count),
				Valid: true,
			},
		}); err != nil {
			return models.OrderID(0), fmt.Errorf("add order item: %w", err)
		}
	}

	return models.OrderID(id), tx.Commit(ctx)
}

func (r *OrdersRepository) GetOrderByID(
	ctx context.Context,
	orderID models.OrderID,
) (*models.Order, error) {

	q := sqlc.New(r.dbpool)
	order, err := q.GetOrderByID(ctx, int64(orderID))
	if err != nil {
		return nil, err
	}
	return &models.Order{
		ID:     orderID,
		UserID: models.UserID(orderID),
		Status: models.GetStatus(order.Status.String),
	}, nil
}

func (r *OrdersRepository) SetStatus(
	ctx context.Context,
	orderID models.OrderID,
	status models.Status,
) error {
	tx, err := r.dbpool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Println(err)
		}
	}()

	q := sqlc.New(r.dbpool)
	q = q.WithTx(tx)
	if err := q.SetStatus(ctx, sqlc.SetStatusParams{
		ID: int64(orderID),
		Status: pgtype.Text{
			String: status.String(),
			Valid:  true,
		},
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
