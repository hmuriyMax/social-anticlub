package pg

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Storage) LoginInfoInsert(ctx context.Context, res *model.LoginResult) error {
	query := `insert into login_info (user_id, token) values ($1, $2)`

	conn, err := s.conn(ctx)
	if err != nil {
		return fmt.Errorf("failed to get connection: %w", err)
	}

	_, err = conn.Exec(ctx, query, res.ID, res.Token)
	return errors.Wrap(err, "Exec")
}
