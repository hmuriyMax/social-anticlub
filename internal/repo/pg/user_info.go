package pg

import (
	"context"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Storage) UserInfoCreate(ctx context.Context, info *model.UserInfo) (userID int64, err error) {
	query := `insert into user_info (first_name, second_name, birthday, gender, hometown, about) values ($1, $2, $3, $4, $5, $6) returning id`

	conn, err := s.conn(ctx)
	if err != nil {
		return 0, err
	}

	err = conn.QueryRow(ctx, query, info.FirstName, info.SecondName, info.Birthday, info.Gender, info.HomeTown, info.About).Scan(&userID)
	return
}
