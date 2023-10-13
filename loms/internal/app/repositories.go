package app

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func getDBConnPool(ctx context.Context) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_DSN"))
	if err != nil {
		return nil, fmt.Errorf("pgxpool: %w", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgxpool ping: %w", err)
	}

	return dbpool, nil
}
