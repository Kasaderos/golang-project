package postgres

import (
	"context"
	"fmt"
	"log"
	"route256/cart/internal/models"
	sqlc "route256/cart/internal/repository/postgres/cart/item"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CartRepository struct {
	dbpool *pgxpool.Pool
}

func NewCartRepostiory(dbpool *pgxpool.Pool) *CartRepository {
	return &CartRepository{
		dbpool: dbpool,
	}
}

func (r *CartRepository) AddItem(
	ctx context.Context,
	userID models.UserID,
	item models.CartItem,
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

	if err := q.AddCartItem(ctx, sqlc.AddCartItemParams{
		UserID: pgtype.Int8{
			Int64: int64(userID),
			Valid: true,
		},
		Sku: pgtype.Int8{
			Int64: int64(item.SKU),
			Valid: true,
		},
		Price: pgtype.Int4{
			Int32: int32(item.Price),
			Valid: true,
		},
		Count: pgtype.Int4{
			Int32: int32(item.Count),
			Valid: true,
		},
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *CartRepository) GetItemsByUserID(
	ctx context.Context,
	userID models.UserID,
) ([]models.CartItem, error) {
	q := sqlc.New(r.dbpool)

	rItems, err := q.GetItemsByUserID(ctx, pgtype.Int8{
		Int64: int64(userID),
		Valid: true,
	})
	if err != nil {
		return nil, fmt.Errorf("get items by user: %w", err)
	}

	items := make([]models.CartItem, 0, len(rItems))
	for _, item := range rItems {
		items = append(items, models.CartItem{
			SKU:   models.SKU(item.Sku.Int64),
			Count: uint16(item.Count.Int32),
			Name:  item.Name.String,
			Price: uint32(item.Price.Int32),
		})
	}

	return items, nil
}

func (r *CartRepository) DeleteItem(
	ctx context.Context,
	userID models.UserID,
	SKU models.SKU,
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

	if err := q.DeleteItem(ctx, sqlc.DeleteItemParams{
		UserID: pgtype.Int8{
			Int64: int64(userID),
			Valid: true,
		},
		Sku: pgtype.Int8{
			Int64: int64(SKU),
		},
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *CartRepository) DeleteItemsByUserID(
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

	if err := q.DeleteItemByUser(ctx, pgtype.Int8{
		Int64: int64(userID),
		Valid: true,
	},
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
