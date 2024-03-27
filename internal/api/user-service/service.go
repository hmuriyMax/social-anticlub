package user_service

import (
	"context"
	"github.com/google/uuid"
	"socialanticlub/internal/pb/user_service"
	"socialanticlub/internal/pkg/users/model"
)

type usersProvider interface {
	Register(ctx context.Context, request *model.RegisterRequest) (*model.RegisterResponse, error)
	Login(ctx context.Context, login uuid.UUID, password string) (*model.LoginResult, error)
}

type Implementation struct {
	user_service.UnimplementedUserServiceServer

	usersProvider usersProvider
}

func NewImplementation(usersProvider usersProvider) *Implementation {
	return &Implementation{
		usersProvider: usersProvider,
	}
}
