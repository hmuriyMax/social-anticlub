package user_service

import (
	"context"
	"github.com/google/uuid"
	"github.com/hmuriyMax/social-anticlub/internal/pb/user_service"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
)

type (
	usersProvider interface {
		GetUserInfo(ctx context.Context, userUUID *uuid.UUID, nick *string) (*model.UserInfo, error)
		Search(ctx context.Context, name, sName string) ([]*model.UserInfo, error)
	}
	authProvider interface {
		Register(ctx context.Context, request *model.RegisterRequest) (*model.RegisterResponse, error)
		Login(ctx context.Context, login uuid.UUID, password string) (*model.TokenInfo, error)
		CheckAuth(userUUID uuid.UUID, tokenString string) error
	}
)

type Implementation struct {
	user_service.UnimplementedUserServiceServer

	usersProvider usersProvider
	authProvider  authProvider
}

func NewImplementation(usersProvider usersProvider, authProvider authProvider) *Implementation {
	return &Implementation{
		usersProvider: usersProvider,
		authProvider:  authProvider,
	}
}
