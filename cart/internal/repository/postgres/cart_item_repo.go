package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"route256/cart/internal/models"
	sqlc "route256/cart/internal/repository/postgres/cart/item"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNoUserItems = errors.New("no user items")

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
		UserID: int64(userID),
		Sku:    int64(item.SKU),
		Price: pgtype.Int4{
			Int32: int32(item.Price),
			Valid: true,
		},
		Amount: pgtype.Int4{
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

	rItems, err := q.GetItemsByUserID(ctx, int64(userID))
	if err != nil {
		return nil, fmt.Errorf("get items by user: %w", err)
	}

	if len(rItems) < 1 {
		return nil, ErrNoUserItems
	}

	items := make([]models.CartItem, 0, len(rItems))
	for _, item := range rItems {
		items = append(items, models.CartItem{
			SKU:   models.SKU(item.Sku),
			Count: uint16(item.Amount.Int32),
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
		UserID: int64(userID),
		Sku:    int64(SKU),
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

	if err := q.DeleteItemByUser(ctx, int64(userID)); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
