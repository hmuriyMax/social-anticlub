package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"socialanticlub/internal/pkg/config"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewClient(ctx context.Context) (*Storage, error) {
	connStr := fmt.Sprintf("user='%s' password='%s' host=%s port=%d dbname=%s pool_max_conns=%d",
		config.GetFromCtx(ctx).PG.User,
		config.GetFromCtx(ctx).PG.Pass,
		config.GetFromCtx(ctx).PG.Host,
		config.GetFromCtx(ctx).PG.Port,
		config.GetFromCtx(ctx).PG.DB,
		config.GetFromCtx(ctx).PG.PoolSize,
	)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping pool: %w", err)
	}

	return &Storage{
		pool: pool,
	}, nil
}
