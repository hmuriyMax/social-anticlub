package pg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Storage) LoginInfoInsert(ctx context.Context, res *model.TokenInfo) error {
	query := `insert into login_info (user_uuid, token) values ($1, $2)`

	conn, err := s.conn(ctx)
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}

	_, err = conn.Exec(ctx, query, res.UserUUID, res.Token)
	return errors.Wrap(err, "Exec")
}

func (s *Storage) LoginInfoSelect(ctx context.Context, token string) (*model.TokenInfo, error) {
	query := `select user_uuid from login_info where token = $1`

	conn, err := s.conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	var userID uuid.UUID
	err = conn.QueryRow(ctx, query, token).Scan(&userID)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return &model.TokenInfo{
		UserUUID: userID,
		Token:    token,
	}, errors.Wrap(err, "QueryRow")
}
