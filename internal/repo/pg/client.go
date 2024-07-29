package pg

import (
	"context"
	"github.com/hmuriyMax/social-anticlub/internal/repo/pg/roles"
	"github.com/jackc/pgx/v5/pgxpool"
)

// dependencies
type (
	pg interface {
		roles.Transaction
		Conn(ctx context.Context, opts ...roles.Option) (*pgxpool.Conn, error)
	}
)

type Storage struct {
	pg pg
	roles.Transaction
}

func NewClient(pg pg) (*Storage, error) {
	return &Storage{
		pg:          pg,
		Transaction: pg,
	}, nil
}
