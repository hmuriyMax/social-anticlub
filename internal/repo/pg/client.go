package pg

import (
	"context"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/config"
)

type Client struct {
	pool *pgx.ConnPool
}

func NewClient(ctx context.Context) (*Client, error) {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     config.GetFromCtx(ctx).PG.Host,
			Port:     config.GetFromCtx(ctx).PG.Port,
			Database: config.GetFromCtx(ctx).PG.DB,
			User:     config.GetFromCtx(ctx).PG.User,
			Password: config.GetFromCtx(ctx).PG.Pass,
		},
		MaxConnections: config.GetFromCtx(ctx).PG.PoolSize,
	})

	return &Client{
			pool: pool,
		},
		errors.Wrap(err, "failed to create pool")
}
