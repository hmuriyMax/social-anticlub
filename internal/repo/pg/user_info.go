package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Storage) UserInfoInsert(ctx context.Context, info *model.UserInfo) (userID int64, err error) {
	query := `insert into user_info (first_name, second_name, birthday, gender, hometown, about) values ($1, $2, $3, $4, $5, $6) returning id`

	conn, err := s.conn(ctx)
	if err != nil {
		return 0, err
	}

	err = conn.QueryRow(ctx, query, info.FirstName, info.SecondName, info.Birthday, info.Gender, info.HomeTown, info.About).Scan(&userID)
	return
}

func (s *Storage) UserInfoSelect(ctx context.Context, userID int64) (info *model.UserInfo, err error) {
	query := `select id, first_name, second_name, birthday, gender, hometown, about from user_info where id = $1`

	conn, err := s.conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	info = &model.UserInfo{}
	err = conn.QueryRow(ctx, query, userID).Scan(&info.ID, &info.FirstName, &info.SecondName, &info.Birthday, &info.Gender, &info.HomeTown, &info.About)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return info, errors.Wrap(err, "QueryRow")
}
