package postgres

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
	sqlc "route256/loms/internal/repository/postgres/stock"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StocksRepository struct {
	dbpool *pgxpool.Pool
}

func NewStocksRepostiory(dbpool *pgxpool.Pool) *StocksRepository {
	return &StocksRepository{
		dbpool: dbpool,
	}
}

func (r *StocksRepository) ReserveStocks(
	ctx context.Context,
	items []models.ItemOrderInfo,
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

	for _, item := range items {
		available, err := q.ReserveStock(ctx, sqlc.ReserveStockParams{
			Available: int64(item.Count),
			Sku:       int64(item.SKU),
		})
		if err != nil {
			return fmt.Errorf("reserve stock: %w, sku=%d, available=%d", err, item.SKU, available)
		}
	}

	return tx.Commit(ctx)
}

func (r *StocksRepository) ReserveRemove(
	ctx context.Context,
	items []models.ItemOrderInfo,
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

	for _, item := range items {
		if err := q.ReserveRemove(ctx, sqlc.ReserveRemoveParams{
			TotalReserved: int64(item.Count),
			Sku:           int64(item.SKU),
		}); err != nil {
			return fmt.Errorf("reserve remove: %w", err)
		}
	}

	return tx.Commit(ctx)
}

func (r *StocksRepository) ReserveCancel(
	ctx context.Context,
	items []models.ItemOrderInfo,
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

	for _, item := range items {
		if err := q.ReserveCancel(ctx, sqlc.ReserveCancelParams{
			Available: int64(item.Count),
			Sku:       int64(item.SKU),
		}); err != nil {
			return fmt.Errorf("reserve remove: %w", err)
		}
	}

	return tx.Commit(ctx)
}

func (r *StocksRepository) GetStocksBySKU(
	ctx context.Context,
	sku models.SKU,
) (uint64, error) {
	tx, err := r.dbpool.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Println(err)
		}
	}()

	q := sqlc.New(r.dbpool)
	q = q.WithTx(tx)

	count, err := q.GetBySKU(ctx, int64(sku))
	if err != nil {
		return 0, err
	}

	return uint64(count), nil
}
