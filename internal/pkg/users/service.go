package users

import (
	"context"
	"github.com/google/uuid"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
	"github.com/hmuriyMax/social-anticlub/internal/repo/pg"
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
