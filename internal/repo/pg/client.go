package pg

import (
	"context"
	"fmt"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewClient(ctx context.Context) (*Storage, error) {
	connStr := fmt.Sprintf("user='%s' password='%s' host=%s port=%d dbname=%s pool_max_conns=%d pool_max_conn_lifetime=%s",
		config.GetFromCtx(ctx).PG.User,
		config.GetFromCtx(ctx).PG.Pass,
		config.GetFromCtx(ctx).PG.Host,
		config.GetFromCtx(ctx).PG.Port,
		config.GetFromCtx(ctx).PG.DB,
		config.GetFromCtx(ctx).PG.PoolSize,
		config.GetFromCtx(ctx).PG.MaxConnLifetime,
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
