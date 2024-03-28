package pg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Storage) UserInfoInsert(ctx context.Context, info *model.UserInfo) (userUUID uuid.UUID, err error) {
	query := `insert into user_info (nickname, first_name, second_name, birthday, gender, hometown, about) values ($1, $2, $3, $4, $5, $6, $7) returning user_uuid`

	conn, err := s.conn(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, query, info.Nickname, info.FirstName, info.SecondName, info.Birthday, info.Gender, info.HomeTown, info.About).Scan(&userUUID)
	return
}

func (s *Storage) UserInfoSelect(ctx context.Context, userUUID *uuid.UUID, nick *string) (info *model.UserInfo, err error) {
	query := `select user_uuid, nickname, first_name, second_name, birthday, gender, hometown, about from user_info`

	if userUUID == nil && nick == nil {
		return nil, nil
	}

	var args []any
	if userUUID != nil {
		query = query + ` where user_uuid = $1`
		args = append(args, *userUUID)
	} else if nick != nil {
		query = query + ` where nickname = $1`
		args = append(args, *nick)
	}

	conn, err := s.conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}
	defer conn.Release()

	info = &model.UserInfo{}
	err = conn.
		QueryRow(ctx, query, args...).
		Scan(&info.UUID, &info.FirstName, &info.FirstName, &info.SecondName, &info.Birthday, &info.Gender, &info.HomeTown, &info.About)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return info, errors.Wrap(err, "QueryRow")
}
