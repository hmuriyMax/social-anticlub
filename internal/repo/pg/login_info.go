package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Storage) LoginInfoInsert(ctx context.Context, res *model.LoginInfo) error {
	query := `insert into login_info (user_id, token) values ($1, $2)`

	conn, err := s.conn(ctx)
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, query, res.ID, res.Token)
	return errors.Wrap(err, "Exec")
}

func (s *Storage) LoginInfoSelect(ctx context.Context, token string) (*model.LoginInfo, error) {
	query := `select user_id from login_info where token = $1`

	conn, err := s.conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}
	defer conn.Release()

	var userID int64
	err = conn.QueryRow(ctx, query, token).Scan(&userID)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return &model.LoginInfo{
		ID:    userID,
		Token: token,
	}, errors.Wrap(err, "QueryRow")
}
