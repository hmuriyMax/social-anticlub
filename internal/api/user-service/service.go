package user_service

import (
	"context"
	"socialanticlub/internal/pb/user_service"
	"socialanticlub/internal/pkg/users/model"
)

type usersProvider interface {
	Register(ctx context.Context, request *model.RegisterRequest) (*model.RegisterResponse, error)
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
