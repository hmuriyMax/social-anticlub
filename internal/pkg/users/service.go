package users

import (
	"context"
	"github.com/google/uuid"
	"socialanticlub/internal/pkg/users/model"
	"socialanticlub/internal/repo/pg"
)

type userRepository interface {
	pg.Connector

	UserInfoSelect(ctx context.Context, userUUID *uuid.UUID, nick *string) (info *model.UserInfo, err error)
	UsersSearch(ctx context.Context, name, sName string) (users []*model.UserInfo, err error)
}

type Service struct {
	repo userRepository
}

func NewService(repo userRepository) *Service {
	return &Service{
		repo: repo,
	}
}
