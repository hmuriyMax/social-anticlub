package pg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Storage) UserAuthInsert(ctx context.Context, auth *model.Login) error {
	query := `insert into user_auth (id, login, pass_hash) values ($1, $2, $3)`

	conn, err := s.conn(ctx)
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}

	_, err = conn.Exec(ctx, query, auth.ID, auth.Login, auth.PassHash)
	return errors.Wrap(err, "Exec")
}

func (s *Storage) UserAuthSelect(ctx context.Context, login uuid.UUID) (*model.Login, error) {
	query := `select id, login, pass_hash from user_auth where login = $1`

	conn, err := s.conn(ctx)
	if err != nil {
		return nil, err
	}

	var auth model.Login
	err = conn.QueryRow(ctx, query, login).Scan(&auth.ID, &auth.Login, &auth.PassHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &auth, nil
}
