package auth

import (
	"context"
	"github.com/google/uuid"
	"socialanticlub/internal/pkg/users/model"
	"socialanticlub/internal/repo/pg"
)

type userRepository interface {
	pg.Connector

	UserInfoInsert(ctx context.Context, info *model.UserInfo) (userUUID uuid.UUID, err error)
	UserInfoSelect(ctx context.Context, userUUID *uuid.UUID, nick *string) (info *model.UserInfo, err error)

	UserAuthInsert(ctx context.Context, auth *model.Login) error
	UserAuthSelect(ctx context.Context, login uuid.UUID) (*model.Login, error)
}

type Service struct {
	repo userRepository
}

func NewService(repo userRepository) *Service {
	return &Service{
		repo: repo,
	}
}
