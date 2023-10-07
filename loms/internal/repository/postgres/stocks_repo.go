package postgres

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
	sqlc "route256/loms/internal/repository/postgres/stock"

	"github.com/jackc/pgx/v5/pgtype"
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
	userID models.UserID,
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
		if err := q.ReserveStock(ctx, sqlc.ReserveStockParams{
			UserID: pgtype.Int8{
				Int64: int64(userID),
				Valid: true,
			},
			Sku: pgtype.Int8{
				Int64: int64(item.SKU),
				Valid: true,
			},
			Count: pgtype.Int8{
				Int64: int64(item.SKU),
				Valid: true,
			},
		}); err != nil {
			return fmt.Errorf("reserver stock: %w", err)
		}
	}

	return tx.Commit(ctx)
}

func (r *StocksRepository) ReserveRemove(
	ctx context.Context,
	userID models.UserID,
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

	stocks, err := q.GetReservedStockByUsedID(ctx, pgtype.Int8{
		Int64: int64(userID),
		Valid: true,
	})
	if err != nil {
		return err
	}

	for _, stock := range stocks {
		if err := q.RemoveStocks(ctx, sqlc.RemoveStocksParams{
			Count: stock.Count,
			Sku:   stock.Sku,
		}); err != nil {
			return fmt.Errorf("remove stocks: %w", err)
		}
	}

	if err := q.DeleteReservedStockByUserID(ctx, pgtype.Int8{
		Int64: int64(userID),
		Valid: true,
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *StocksRepository) ReserveCancel(
	ctx context.Context,
	userID models.UserID,
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

	if err := q.DeleteReservedStockByUserID(ctx, pgtype.Int8{
		Int64: int64(userID),
		Valid: true,
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *StocksRepository) GetStocksBySKU(
	ctx context.Context,
	SKU models.SKU,
) (count uint64, err error) {
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

	stocksCount, err := q.CountStocksBySKU(ctx, pgtype.Int8{
		Int64: int64(SKU),
	})
	if err != nil {
		return 0, err
	}

	reservedStock, err := q.CountReservedStocksBySKU(ctx, pgtype.Int8{
		Int64: int64(SKU),
	})
	if err != nil {
		return 0, err
	}

	return uint64(stocksCount.Int64 - reservedStock), tx.Commit(ctx)
}
