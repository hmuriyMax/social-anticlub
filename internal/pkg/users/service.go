package users

import (
	"context"
	"socialanticlub/internal/pkg/users/model"
	"socialanticlub/internal/repo/pg"
)

type userRepository interface {
	pg.Connector
	UserInfoCreate(ctx context.Context, info *model.UserInfo) (userID int64, err error)
	UserAuthInsert(ctx context.Context, auth *model.Login) error
}

type Service struct {
	repo userRepository
}

func NewService(repo userRepository) *Service {
	return &Service{
		repo: repo,
	}
}
