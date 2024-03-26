package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	RoleRead  = "read"
	RoleWrite = "write"
)

type ctxConnKey struct{}

type Connector interface {
	ExecTx(ctx context.Context, action func(ctx context.Context) error) error
}

func (s *Storage) ExecTx(ctx context.Context, action func(ctx context.Context) error) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	actionCtx := context.WithValue(ctx, ctxConnKey{}, tx.Conn())
	err = action(actionCtx)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return fmt.Errorf("failed to rollback transaction: %w", rollbackErr)
		}
		return err
	}
	return errors.Wrap(tx.Commit(ctx), "failed to commit transaction")
}

func (s *Storage) conn(ctx context.Context) (*pgx.Conn, error) {
	conn, ok := ctx.Value(ctxConnKey{}).(*pgx.Conn)
	if !ok {
		poolConn, err := s.pool.Acquire(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to acquire connection from pool: %w", err)
		}
		return poolConn.Conn(), nil

	}
	return conn, nil
}
