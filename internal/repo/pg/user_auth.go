package pg

import (
	"context"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Storage) UserAuthInsert(ctx context.Context, auth *model.Login) error {
	query := `insert into user_auth (id, login, pass_hash) values ($1, $2, $3)`

	conn, err := s.conn(ctx)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, query, auth.ID, auth.Login, auth.PassHash)
	return err
}
