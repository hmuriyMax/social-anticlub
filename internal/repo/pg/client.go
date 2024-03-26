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
	//pgConfig := &pgxpool.Config{
	//	ConnConfig: &pgx.ConnConfig{
	//		Config: pgconn.Config{
	//			Host:     config.GetFromCtx(ctx).PG.Host,
	//			Port:     config.GetFromCtx(ctx).PG.Port,
	//			Database: config.GetFromCtx(ctx).PG.DB,
	//			User:     config.GetFromCtx(ctx).PG.User,
	//			Password: config.GetFromCtx(ctx).PG.Pass,
	//		},
	//	},
	//	MaxConns: config.GetFromCtx(ctx).PG.PoolSize,
	//}

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

	return &Storage{
		pool: pool,
	}, pool.Ping(ctx)
}
