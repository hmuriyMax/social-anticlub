package roles

import (
	"context"
	"fmt"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PG struct {
	master       *pgxpool.Pool
	asyncReplica *pgxpool.Pool

	defaultRole *ConnRole
}

func NewPG(ctx context.Context, cnf *config.Config) (*PG, error) {
	// master connect
	connStr := fmt.Sprintf("user='%s' password='%s' host=%s port=%d dbname=%s pool_max_conns=%d pool_max_conn_lifetime=%s",
		cnf.PG.Master.User,
		cnf.PG.Master.Pass,
		cnf.PG.Master.Host,
		cnf.PG.Master.Port,
		cnf.PG.Master.DB,
		cnf.PG.Master.PoolSize,
		cnf.PG.Master.MaxConnLifetime,
	)

	masterPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}
	if err = masterPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping pool: %w", err)
	}

	// async replica connect
	connStr = fmt.Sprintf("user='%s' password='%s' host=%s port=%d dbname=%s pool_max_conns=%d pool_max_conn_lifetime=%s",
		cnf.PG.Async.User,
		cnf.PG.Async.Pass,
		cnf.PG.Async.Host,
		cnf.PG.Async.Port,
		cnf.PG.Async.DB,
		cnf.PG.Async.PoolSize,
		cnf.PG.Async.MaxConnLifetime,
	)

	asyncPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}
	if err = asyncPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping pool: %w", err)
	}

	var defaultRoute *ConnRole
	switch cnf.PG.DefaultRole {
	case "master", "write", "main", "":
		defaultRoute = RoleMaster
	case "read", "async", "slave":
		defaultRoute = RoleAsync
	default:
		return nil, fmt.Errorf("invalid default role: %s", cnf.PG.DefaultRole)
	}

	return &PG{
		master:       masterPool,
		asyncReplica: asyncPool,
		defaultRole:  defaultRoute,
	}, nil
}
