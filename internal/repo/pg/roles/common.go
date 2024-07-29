package roles

import (
	"context"
	"fmt"
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type ConnRole string

var (
	RoleAsync  = helpers.Ptr(ConnRole("async"))
	RoleMaster = helpers.Ptr(ConnRole("master"))
)

type ctxConnKey struct{}

type Transaction interface {
	ExecTx(ctx context.Context, action func(ctx context.Context) error) error
}

func (s *PG) ExecTx(ctx context.Context, action func(ctx context.Context) error) error {
	conn, err := s.Conn(ctx)
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
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

// TODO: вынести в отдельный файл
type (
	connParams struct {
		role *ConnRole
	}

	Option func(*connParams)
)

func (s *PG) Conn(ctx context.Context, opts ...Option) (*pgxpool.Conn, error) {
	var params = &connParams{}
	for _, opt := range opts {
		opt(params)
	}

	if params.role == nil {
		params.role = s.defaultRole
	}

	var pool *pgxpool.Pool
	switch params.role {
	case RoleMaster:
		pool = s.master
	default:
		pool = s.asyncReplica
	}

	conn, ok := ctx.Value(ctxConnKey{}).(*pgxpool.Conn)
	if !ok {
		poolConn, err := pool.Acquire(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to acquire connection from pool: %w", err)
		}
		return poolConn, nil

	}
	return conn, nil
}

func WithRole(role *ConnRole) Option {
	return func(params *connParams) {
		params.role = role
	}
}
