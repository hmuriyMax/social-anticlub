package user_service

import (
	"socialanticlub/internal/pb/user_service"
)

type usersProvider interface {
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
