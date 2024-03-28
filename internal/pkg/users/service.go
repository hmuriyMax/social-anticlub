package users

import (
	"context"
	"github.com/google/uuid"
	"socialanticlub/internal/pkg/users/model"
	"socialanticlub/internal/repo/pg"
)

type userRepository interface {
	pg.Connector

	UserInfoInsert(ctx context.Context, info *model.UserInfo) (userUUID uuid.UUID, err error)
	UserInfoSelect(ctx context.Context, userUUID uuid.UUID) (info *model.UserInfo, err error)

	UserAuthInsert(ctx context.Context, auth *model.Login) error
	UserAuthSelect(ctx context.Context, login uuid.UUID) (*model.Login, error)

	LoginInfoInsert(ctx context.Context, res *model.TokenInfo) error
	LoginInfoSelect(ctx context.Context, token string) (*model.TokenInfo, error)
}

type Service struct {
	repo userRepository
}

func NewService(repo userRepository) *Service {
	return &Service{
		repo: repo,
	}
}
