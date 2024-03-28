package user_service

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"socialanticlub/internal/pb/user_service"
	"socialanticlub/internal/pkg/users/model"
)

func (i *Implementation) Auth(ctx context.Context, req *user_service.AuthRequest) (*user_service.AuthResponse, error) {
	if err := validateAuth(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	login, err := uuid.Parse(req.GetLogin())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	loginInfo, err := i.usersProvider.Login(ctx, login, req.GetPass())
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNoUser):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, model.ErrWrongPassword):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &user_service.AuthResponse{
		User: &user_service.LoginInfo{
			UserID: loginInfo.UserUUID.String(),
			Token:  loginInfo.Token,
		},
	}, nil
}

func validateAuth(req *user_service.AuthRequest) error {
	switch {
	case req == nil:
		return errors.New("auth info required")
	case req.Login == "":
		return errors.New("auth login required")
	case req.Pass == "":
		return errors.New("auth password required")
	}
	return nil
}
