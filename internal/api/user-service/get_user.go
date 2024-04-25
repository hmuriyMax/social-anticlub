package user_service

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"socialanticlub/internal/api/user-service/converters"
	"socialanticlub/internal/helpers"
	"socialanticlub/internal/pb/user_service"
	"socialanticlub/internal/pkg/users/model"
)

func (i *Implementation) GetUser(ctx context.Context, req *user_service.GetUserRequest) (*user_service.GetUserResponse, error) {
	if err := validateGetUser(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	loginInfo := helpers.GetAuthInfo(ctx)
	if err := i.authProvider.CheckAuth(loginInfo.GetUUID(), loginInfo.GetToken()); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var (
		nick     *string
		userUUID *uuid.UUID
	)
	parsedUUID, err := uuid.Parse(req.GetIdentifier())
	if err != nil {
		nick = helpers.Ptr(req.GetIdentifier())
	} else {
		userUUID = &parsedUUID
	}

	userInfo, err := i.usersProvider.GetUserInfo(ctx, userUUID, nick)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNoUser):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &user_service.GetUserResponse{
		User:    converters.UserInfoToPB(userInfo),
		IsOwner: !userInfo.Birthday.IsZero(),
	}, nil

}

func validateGetUser(req *user_service.GetUserRequest) error {
	switch {
	case req == nil:
		return errors.New("user id required")
	case len(req.GetIdentifier()) == 0:
		return errors.New("user id required")
	}
	return nil
}
